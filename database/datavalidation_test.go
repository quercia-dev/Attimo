package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
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
			assert.Equal(t, tt.want, validateURL(tt.input), "validateURL(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateInRange(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		min   int
		max   int
		want  bool
	}{
		{"Valid - in range", 3, 1, 5, true},
		{"Invalid - below range", 0, 1, 5, false},
		{"Invalid - above range", 6, 1, 5, false},
		{"Invalid - wrong type", "3", 1, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateInRange(tt.input, tt.min, tt.max), "validateInRange(%v, %v, %v) should return %v", tt.input, tt.min, tt.max, tt.want)
		})
	}
}

func TestValidateEmail(t *testing.T) {
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
			assert.Equal(t, tt.want, validateEmail(tt.input), "validateEmail(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidatePhone(t *testing.T) {
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
			assert.Equal(t, tt.want, validatePhone(tt.input), "validatePhone(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateFileExists(t *testing.T) {
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
			assert.Equal(t, tt.want, validateFileExists(tt.input), "validateFileExists(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestValidateDate(t *testing.T) {
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
			assert.Equal(t, tt.want, validateDate(tt.input), "validateDate(%v) should return %v", tt.input, tt.want)
		})
	}
}

func TestDatatype_ValidateValue(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "example")
	assert.NoError(t, err, "Failed to create temporary file")
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name       string
		valueCheck string
		input      interface{}
		want       bool
	}{
		{"URL check", "URL", "https://example.com", true},
		{"Range check", "in{1,2,3,4,5}", 3, true},
		{"Email check", "mail_ping", "test@example.com", true},
		{"Phone check", "phone", "1234567890", true},
		{"File check", "file_exists", tmpfile.Name(), true},
		{"Date check", "date", "01-01-2024", true},
		{"Invalid check type", "invalid_check", "any value", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &Datatype{ValueCheck: tt.valueCheck}
			assert.Equal(t, tt.want, dt.ValidateValue(tt.input), "Datatype.ValidateValue(%v) with ValueCheck=%v should return %v", tt.input, tt.valueCheck, tt.want)
		})
	}
}
