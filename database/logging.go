package database

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogging(logDir string) error {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil { //0755 is the permissions for mkdir
		return fmt.Errorf("failed to create log dir: %w", err)
	}

	// open log file
	logFile, err := os.OpenFile(filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create multi-writers to file/terminal
	infoWriter := io.MultiWriter(os.Stdout, logFile)
	warningWriter := io.MultiWriter(os.Stdout, logFile)
	errorWriter := io.MultiWriter(os.Stdout, logFile)

	InfoLogger = log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(warningWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// LogError logs an error and returns it wrapped with additional context
func LogError(context string, err error) error {
	if err != nil {
		wrappedErr := fmt.Errorf("%s: %w", context, err)
		ErrorLogger.Println(wrappedErr)
		return wrappedErr
	}
	return nil
}

// logTypeError logs a warning message when a value is not of the expected type
// It uses the global WarningLogger if it is not nil, otherwise it uses fmt.Printf
// The message is in the format: "%v needs to be a %s"
func logTypeError(value interface{}, typeName string) {
	if WarningLogger != nil {
		WarningLogger.Printf("%v needs to be a %s", value, typeName)
	} else {
		fmt.Printf("%v needs to be a %s", value, typeName)
	}
}

// logArgsError logs a warning message when the number of arguments is not as expected
// It uses the global WarningLogger if it is not nil, otherwise it uses fmt.Printf
// The message is in the format: "Expected %d arguments, got %d: %v"
func logArgsError(args []string, expected int) {
	if WarningLogger != nil {
		WarningLogger.Printf("Expected %d arguments, got %d: %v", expected, len(args), args)
	} else {
		fmt.Printf("Expected %d arguments, got %d: %v", expected, len(args), args)
	}
}
