package database

import (
	"net/mail"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// SplitType splits the type into the type and a list of string parameters
// "IN(1,2,4,4)" -> "IN", ["1", "2", "3", "4"]
// Uses the first open parenthesis to split the string,
// and then splits the parameters by commas until the penultimate character
func SplitStringArgument(input string) (string, []string) {
	openParenIndex := strings.Index(input, "(")
	if openParenIndex == -1 {
		// No parameters
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

// / ValueCheck switch cases
func (dt *Datatype) ValidateCheck(value interface{}) bool {

	typeS, args := SplitStringArgument(dt.ValueCheck)

	switch typeS {
	case nonemptyCheck:
		return validateNonempty(value)
	case RangeCheck:
		return validateInRange(value, args)
	case SetCheck:
		return validateSet(value, args)
	case NoCheck:
		return true
	case URLCheck:
		return validateURL(value)
	case MailCheck:
		return validateEmail(value)
	case PhoneCheck:
		return validatePhone(value)
	case FileCheck:
		return validateFileExists(value)
	case DateCheck:
		return validateDate(value)
	default:
		WarningLogger.Printf("Unknown validation type: %s", typeS)
		return false
	}
}

func validateNonempty(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		logTypeError(value, "string")
		return false
	}
	return str != ""
}

func validateSet(value interface{}, args []string) bool {
	str, ok := value.(string)
	if !ok {
		logTypeError(value, "string")
		return false
	}
	for _, arg := range args {
		if str == arg {
			return true
		}
	}
	return false
}

// /rejects empty http:// and relative urls like /foo/bar
func validateURL(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		logTypeError(value, "string")
		return false
	}
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateInRange(value interface{}, args []string) bool {
	if args == nil || len(args) != 2 {
		logArgsError(args, 2)
		return false
	}
	i, ok := value.(int)
	if !ok {
		logTypeError(value, "int")
		return false
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		WarningLogger.Printf("Failed to parse min value: %v", err)
		return false
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		WarningLogger.Printf("Failed to parse max value: %v", err)
		return false
	}
	return i >= min && i <= max
}

func validateEmail(value interface{}) bool {
	email, ok := value.(string)
	if !ok {
		logTypeError(value, "string")
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePhone(value interface{}) bool {
	number, ok := value.(string)
	if !ok {
		logTypeError(value, "string")
		return false
	}
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(number) && len(number) >= 7 && len(number) <= 15
}

func validateFileExists(value interface{}) bool {
	path, ok := value.(string)

	if !ok {
		logTypeError(value, "string")
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

func validateDate(value interface{}) bool {
	date, ok := value.(string)
	if !ok {
		logTypeError(value, "string")
		return false
	}
	layout := "02-01-2006"
	_, err := time.Parse(layout, date)
	return err == nil
}
