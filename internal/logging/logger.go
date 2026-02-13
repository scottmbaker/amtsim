package logging

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	LEVEL_ERROR LogLevel = iota
	LEVEL_INFO
	LEVEL_DEBUG
)

var (
	currentLevel = LEVEL_INFO
	logger       = log.New(os.Stderr, "", log.LstdFlags)
)

func SetLevel(level LogLevel) {
	currentLevel = level
}

func Debug(format string, v ...interface{}) {
	if currentLevel >= LEVEL_DEBUG {
		_ = logger.Output(2, fmt.Sprintf("[DEBUG] "+format, v...))
	}
}

func Info(format string, v ...interface{}) {
	if currentLevel >= LEVEL_INFO {
		_ = logger.Output(2, fmt.Sprintf("[INFO]  "+format, v...))
	}
}

func Error(format string, v ...interface{}) {
	if currentLevel >= LEVEL_ERROR {
		_ = logger.Output(2, fmt.Sprintf("[ERROR] "+format, v...))
	}
}

func Fatal(format string, v ...interface{}) {
	_ = logger.Output(2, fmt.Sprintf("[FATAL] "+format, v...))
	os.Exit(1)
}
