package database

import (
	"os"
	"testing"

	log "Attimo/logging"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"Valid URL", "https://example.com", true},
		{"Invalid URL - relative path", "/foo/bar", false},
		{"Invalid URL - empty scheme", "example.com", false},
		{"Invalid URL - wrong type", 42, false},
		{"Invalid URL - empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateURL(logger, tt.input), "validateURL(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateInRange(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tests := []struct {
		name  string
		input interface{}
		args  []string
		want  bool
	}{
		{"Valid - in range", 3, []string{"1", "5"}, true},
		{"Valid - at lower bound", 1, []string{"1", "5"}, true},
		{"Valid - at upper bound", 5, []string{"1", "5"}, true},
		{"Valid - negative range", -3, []string{"-5", "-1"}, true},
		{"Valid - single value", 3, []string{"3", "3"}, true},

		{"Invalid - empty args", 3, []string{}, false},
		{"Invalid - lower number of args", 3, []string{"1"}, false},
		{"Invalid - higher number of args", 3, []string{"1", "2", "3"}, false},

		{"Invalid - below range", 0, []string{"1", "10"}, false},
		{"Invalid - above range", 6, []string{"-1", "5"}, false},
		{"Invalid - wrong type", "3", []string{"1", "3"}, false},
		{"Invalid - missing arguments", 3, []string{"1"}, false},
		{"Invalid - empty string", "", []string{"1", "3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateInRange(logger, tt.input, tt.args), "validateInRange(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateInSet(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tests := []struct {
		name  string
		input interface{}
		args  []string
		want  bool
	}{
		{"Valid - in set", "foo", []string{"foo", "bar", "baz"}, true},
		{"Valid - first in set", "foo", []string{"foo", "bar", "baz"}, true},
		{"Valid - last in set", "baz", []string{"foo", "bar", "baz"}, true},
		{"Valid - single value", "foo", []string{"foo"}, true},
		{"Valid - empty set", "foo", []string{}, false},
		{"Valid - empty string", "", []string{"foo", "bar", "baz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateInSet(logger, tt.input, tt.args), "validateInSet(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateEmail(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"Valid email", "test@example.com", true},
		{"Invalid email - no @", "testexample.com", false},
		{"Invalid email - wrong type", 42, false},
		{"Invalid email - empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateEmail(logger, tt.input), "validateEmail(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidatePhone(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"Valid phone - 7 digits", "1234567", true},
		{"Valid phone - 15 digits", "123456789012345", true},
		{"Invalid phone - too short", "123456", false},
		{"Invalid phone - too long", "1234567890123456", false},
		{"Invalid phone - non-numeric", "123abc456", false},
		{"Invalid phone - wrong type", 1234567, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validatePhone(logger, tt.input), "validatePhone(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateFileExists(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "example")
	assert.NoError(t, err, "Failed to create temporary file")
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"Valid file", tmpfile.Name(), true},
		{"Invalid file - doesn't exist", "/path/to/nonexistent/file", false},
		{"Invalid file - wrong type", 42, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateFileExists(logger, tt.input), "validateFileExists(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateDate(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"Valid date", "01-01-2024", true},
		{"Invalid date - wrong format", "2024-01-01", false},
		{"Invalid date - wrong type", 42, false},
		{"Invalid date - empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateDate(logger, tt.input), "validateDate(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateCheck(t *testing.T) {
	logger, err := log.GetTestLogger()
	if err != nil {
		t.Errorf(log.LoggerErrorString, err)
	}

	tmpfile, err := os.CreateTemp("", "example")
	assert.NoError(t, err, "Failed to create temporary file")
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name       string
		valueCheck string
		input      interface{}
		want       bool
	}{
		{"URL check", URLCheck, "https://example.com", true},
		{"Set check", SetCheck + "(ABS,Sd,sdasd,sad asd,2312)", "2312", true},
		{"Range check", RangeCheck + "(1,5)", 3, true},
		{"Email check", MailCheck, "test@example.com", true},
		{"Phone check", PhoneCheck, "1234567890", true},
		{"File check", FileCheck, tmpfile.Name(), true},
		{"Date check", DateCheck, "01-01-2024", true},
		{"Invalid check type", "invalid_check", "any value", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &Datatype{ValueCheck: tt.valueCheck}
			assert.Equal(t, tt.want, dt.ValidateCheck(logger, tt.input), "Datatype.ValidateCheck(%v) with ValueCheck=%v should return %v", tt.input, tt.valueCheck, tt.want)
		})
	}
}
