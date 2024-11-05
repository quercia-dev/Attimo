package database

import (
	"Attimo/logging"
	"database/sql"
	"fmt"
	"net/mail"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	TypeMismatch = " %v is not %s"
)

// SplitStringArgument splits the type into the type and a list of string parameters
// "IN(1,2,3,4)" -> "IN", ["1", "2", "3", "4"]
func SplitStringArgument(input string) (string, []string) {
	openParenIndex := strings.Index(input, "(")
	if openParenIndex == -1 {
		return input, nil
	}

	typeName := strings.TrimSpace(input[:openParenIndex])
	paramsString := input[openParenIndex+1 : len(input)-1]
	params := strings.Split(paramsString, ",")

	for i, param := range params {
		params[i] = strings.TrimSpace(param)
	}

	return typeName, params
}

// TODO REMOVE LOGGING WHEN OPERATIONS BECOME MORE FREQUENT
// ValidateCheck performs validation based on the datatype's check rules
func (dt *Datatype) ValidateCheck(value interface{}, logger *logging.Logger) bool {
	typeS, args := SplitStringArgument(dt.ValueCheck)

	switch typeS {
	case nonemptyCheck:
		valid := validateNonempty(value)
		if !valid {
			logger.LogWarn("Nonempty validation failed for value: %v", value)
		}
		return valid
	case RangeCheck:
		valid := validateInRange(value, args)
		if !valid {
			logger.LogWarn("Range validation failed for value: %v (args: %v)", value, args)
		}
		return valid
	case SetCheck:
		valid := validateInSet(value, args)
		if !valid {
			logger.LogWarn("Set validation failed for value: %v (args: %v)", value, args)
		}
		return valid
	case NoCheck:
		return true
	case URLCheck:
		valid := validateURL(value)
		if !valid {
			logger.LogWarn("URL validation failed for value: %v", value)
		}
		return valid
	case MailCheck:
		valid := validateEmail(value)
		if !valid {
			logger.LogWarn("Email validation failed for value: %v", value)
		}
		return valid
	case PhoneCheck:
		valid := validatePhone(value)
		if !valid {
			logger.LogWarn("Phone validation failed for value: %v", value)
		}
		return valid
	case FileCheck:
		valid := validateFileExists(value)
		if !valid {
			logger.LogWarn("File validation failed for value: %v", value)
		}
		return valid
	case DateCheck:
		valid := validateDate(value)
		if !valid {
			logger.LogWarn("Date validation failed for value: %v", value)
		}
		return valid
	default:
		logger.LogErr("Unrecognized type: %v", dt.ValueCheck)
		return false
	}
}

// Individual validation functions
func validateNonempty(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		return false
	}
	return str != ""
}

func validateInSet(value interface{}, args []string) bool {
	str, ok := value.(string)
	if !ok {
		return false
	}
	for _, arg := range args {
		if str == arg {
			return true
		}
	}
	return false
}

func validateURL(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		return false
	}
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateInRange(value interface{}, args []string) bool {
	if args == nil || len(args) != 2 {
		return false
	}
	i, ok := value.(int)
	if !ok {
		return false
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		return false
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

func validateEmail(value interface{}) bool {
	email, ok := value.(string)
	if !ok {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePhone(value interface{}) bool {
	number, ok := value.(string)
	if !ok {
		return false
	}
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(number) && len(number) >= 7 && len(number) <= 15
}

func validateFileExists(value interface{}) bool {
	path, ok := value.(string)
	if !ok {
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

func validateDate(value interface{}) bool {
	date, ok := value.(string)
	if !ok {
		return false
	}
	_, err := time.Parse(dateFormat, date)
	return err == nil
}

// getDatatypeByName retrieves a datatype from the database by name
func getDatatypeByName(tx *sql.Tx, name string) (*Datatype, error) {
	var dt Datatype
	err := tx.QueryRow(`
		SELECT id, name, variable_type, completion_value, completion_sort, value_check 
		FROM datatypes 
		WHERE name = ?
	`, name).Scan(&dt.ID, &dt.Name, &dt.VariableType, &dt.CompletionValue, &dt.CompletionSort, &dt.ValueCheck)

	if err != nil {
		return nil, fmt.Errorf("failed to get datatype %s: %w", name, err)
	}
	return &dt, nil
}

// validateField validates a single field value against its datatype
func (d *Database) validateField(tx *sql.Tx, columnName string, value interface{}) error {
	datatype, err := getDatatypeByName(tx, columnName)
	if err != nil {
		return fmt.Errorf("failed to get datatype for column %s: %w", columnName, err)
	}

	if !datatype.ValidateCheck(value, d.logger) {
		return fmt.Errorf("invalid value for column %s: %v", columnName, value)
	}

	return nil
}

// validateInputData validates all fields in the input data
func (d *Database) validateInputData(tx *sql.Tx, categoryName string, data map[string]interface{}) error {
	// Get column information
	rows, err := tx.Query(`
		SELECT name 
		FROM pragma_table_info(?)
		WHERE name != 'id' 
		  AND name != 'created_at' 
		  AND name != 'updated_at' 
		  AND name != 'deleted_at'
	`, categoryName)
	if err != nil {
		return fmt.Errorf("failed to get column info: %w", err)
	}
	defer rows.Close()

	// Create a map of valid columns
	validColumns := make(map[string]bool)
	for rows.Next() {
		var colName string
		if err := rows.Scan(&colName); err != nil {
			return fmt.Errorf("error scanning column name: %w", err)
		}
		validColumns[colName] = true
	}

	// Validate each provided field
	for field, value := range data {
		if !validColumns[field] {
			return fmt.Errorf("invalid field name: %s", field)
		}

		if err := d.validateField(tx, field, value); err != nil {
			return err
		}
	}

	return nil
}
