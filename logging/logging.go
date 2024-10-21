package logging

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

func InitLoggingWithWriter(w io.Writer) error {
	if w == nil {
		return fmt.Errorf("writer cannot be nil")
	}

	InfoLogger = log.New(w, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(w, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(w, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

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

	infoWriter := io.Writer(logFile)
	warningWriter := io.Writer(logFile)
	errorWriter := io.Writer(logFile)

	InfoLogger = log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(warningWriter, "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime)

	return nil
}

func LogInfo(format string, v ...interface{}) error {
	if InfoLogger != nil {
		InfoLogger.Printf(format, v...)
		return nil
	} else {
		_, err := fmt.Printf(format+"\n", v...)
		return err
	}
}

func LogWarn(format string, v ...interface{}) error {
	if WarningLogger != nil {
		WarningLogger.Printf(format, v...)
		return nil
	} else {
		_, err := fmt.Printf(format+"\n", v...)
		return err
	}
}

func LogErr(format string, v ...interface{}) error {
	if ErrorLogger != nil {
		ErrorLogger.Printf(format, v...)
		return nil
	} else {
		_, err := fmt.Printf(format+"\n", v...)
		return err
	}
}
