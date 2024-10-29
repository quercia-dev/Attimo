package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	logfile           = "logs/app.log"
	LoggerErrorString = "Could not create logger: %v"
)

type Logger struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

func GetTestLogger() (*Logger, error) {
	return InitLogging(logfile)
}

func InitLoggingWithWriter(w io.Writer) (*Logger, error) {
	if w == nil {
		return nil, fmt.Errorf("writer cannot be nil")
	}

	return &Logger{
		InfoLogger:    log.New(w, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarningLogger: log.New(w, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   log.New(w, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}

func InitLogging(logDir string) (*Logger, error) {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil { //0755 is the permissions for mkdir
		return nil, fmt.Errorf("failed to create log dir: %w", err)
	}

	// open log file
	logFile, err := os.OpenFile(filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	infoWriter := io.Writer(logFile)
	warningWriter := io.Writer(logFile)
	errorWriter := io.Writer(logFile)

	return &Logger{
		InfoLogger:    log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime),
		WarningLogger: log.New(warningWriter, "WARNING: ", log.Ldate|log.Ltime),
		ErrorLogger:   log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime),
	}, nil
}

func (log *Logger) LogInfo(format string, v ...interface{}) error {
	if log.InfoLogger != nil {
		log.InfoLogger.Printf(format, v...)
		return nil
	} else {
		_, err := fmt.Printf(format+"\n", v...)
		return err
	}
}

func (log *Logger) LogWarn(format string, v ...interface{}) error {
	if log.WarningLogger != nil {
		log.WarningLogger.Printf(format, v...)
		return nil
	} else {
		_, err := fmt.Printf(format+"\n", v...)
		return err
	}
}

func (log *Logger) LogErr(format string, v ...interface{}) error {
	if log.ErrorLogger != nil {
		log.ErrorLogger.Printf(format, v...)
		return nil
	} else {
		_, err := fmt.Printf(format+"\n", v...)
		return err
	}
}
