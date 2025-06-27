package session

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Session struct {
	logFilePath string
	logFileName string
}

func New() Session {
	// Setup logging file path and directory at the start.
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Fprintf(os.Stderr, "critical: unable to find user home directory: %v\n", err)
		os.Exit(1)
	}
	logDir := filepath.Join(homeDir, ".config", "funkeytype", "logs")

	// Create the log directory if it doesn't exist.
	// 0755 gives rwx for owner, rx for group and others.
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "critical: unable to create log directory %s: %v\n", logDir, err)
		os.Exit(1)
	}

	// Format the log file name with a timestamp.
	logFileName := "funkeytype-log-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
	logFilePath := filepath.Join(logDir, logFileName)

	newSession := Session{
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	return newSession
}

func (s *Session) LogErrorAndExit(err error) {
	// This function will be called when a fatal error occurs.
	// It logs the error to the designated log file and terminates the program.
	f, fileErr := os.OpenFile(s.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		// If we can't even open the log file, something is seriously wrong.
		// We'll print both the original error and the file opening error to stderr.
		fmt.Fprintf(os.Stderr, "FATAL: could not open log file %s: %v. Original error was: %v\n", s.logFilePath, fileErr, err)
		os.Exit(1)
	}
	defer f.Close()

	// Prepend the timestamp to the log message
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Add file/line number
	log.Output(2, fmt.Sprintf("FATAL: %v", err)) // log the error
	os.Exit(1)                                   // exit
}
