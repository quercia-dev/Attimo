package database

import (
	log "Attimo/logging"
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
func (dt *Datatype) ValidateCheck(logger *log.Logger, value interface{}) bool {
	if logger == nil {
		fmt.Println("nil logger")
		return false
	}

	typeS, args := SplitStringArgument(dt.ValueCheck)

	switch typeS {
	case nonemptyCheck:
		return validateNonempty(logger, value)
	case RangeCheck:
		return validateInRange(logger, value, args)
	case SetCheck:
		return validateInSet(logger, value, args)
	case NoCheck:
		return true
	case URLCheck:
		return validateURL(logger, value)
	case MailCheck:
		return validateEmail(logger, value)
	case PhoneCheck:
		return validatePhone(logger, value)
	case FileCheck:
		return validateFileExists(logger, value)
	case DateCheck:
		return validateDate(logger, value)
	default:
		logger.LogErr("Unrecognized type: %v", dt.ValueCheck)
		return false
	}
}

func validateNonempty(logger *log.Logger, value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		logger.LogInfo(TypeMismatch, value, "string")
		return false
	}
	return str != ""
}

func validateInSet(logger *log.Logger, value interface{}, args []string) bool {
	str, ok := value.(string)
	if !ok {
		logger.LogInfo(TypeMismatch, value, "string")
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
func validateURL(logger *log.Logger, value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		logger.LogInfo(" %v is not a string", value)
		return false
	}
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateInRange(logger *log.Logger, value interface{}, args []string) bool {
	if args == nil || len(args) != 2 {
		logger.LogInfo("args are not exactly 2 in length: %v", args)
		return false
	}
	i, ok := value.(int)
	if !ok {
		logger.LogInfo(TypeMismatch, value, "int")
		return false
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		logger.LogInfo(TypeMismatch, args[0], "int")
		return false
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		logger.LogInfo(TypeMismatch, args[1], "int")
		return false
	}
	return i >= min && i <= max
}

func validateEmail(logger *log.Logger, value interface{}) bool {
	email, ok := value.(string)
	if !ok {
		logger.LogInfo(TypeMismatch, value, "string")
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePhone(logger *log.Logger, value interface{}) bool {
	number, ok := value.(string)
	if !ok {
		logger.LogInfo(TypeMismatch, value, "string")
		return false
	}
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(number) && len(number) >= 7 && len(number) <= 15
}

func validateFileExists(logger *log.Logger, value interface{}) bool {
	path, ok := value.(string)

	if !ok {
		logger.LogInfo(TypeMismatch, value, "string")
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

func validateDate(logger *log.Logger, value interface{}) bool {
	date, ok := value.(string)
	if !ok {
		logger.LogInfo(TypeMismatch, value, "string")
		return false
	}
	layout := dateFormat
	_, err := time.Parse(layout, date)
	return err == nil
}
