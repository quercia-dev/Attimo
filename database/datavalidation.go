package database

import (
	"net/mail"
	"net/url"
	"os"
	"regexp"
	"time"
)

func logTypeError(value interface{}, typeName string) {
	WarningLogger.Printf("%v needs to be a %s", value, typeName)
}

// / ValueCheck switch cases
func (dt *Datatype) ValidateValue(value interface{}) bool {
	switch dt.ValueCheck {
	case "URL":
		return validateURL(value)
	case "in{1,2,3,4,5}": // TEMP
		return validateInRange(value, 1, 5)
	case "mail_ping":
		return validateEmail(value)
	case "phone":
		return validatePhone(value)
	case "file_exists":
		return validateFileExists(value)
	case "date":
		return validateDate(value)
	default:
		return false
	}
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

func validateInRange(value interface{}, min, max int) bool {
	i, ok := value.(int) //assert value is an integer
	if !ok {
		logTypeError(value, "int")
		return false
	}
	return ok && i >= min && i <= max
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
