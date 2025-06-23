package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu     sync.Mutex
	logger *log.Logger
	level  LogLevel
}

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

var logColors = map[LogLevel]string{
	DebugLevel: ColorBlue,
	InfoLevel:  ColorGreen,
	WarnLevel:  ColorYellow,
	ErrorLevel: ColorRed,
}

var logLevels = map[LogLevel]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
}

var instance *Logger
var once sync.Once

// InitLogger initializes the global logger instance
func InitLogger(level LogLevel, output string) (*Logger, error) {
	var err error
	once.Do(func() {
		var out *os.File
		if output == "stdout" || output == "" {
			out = os.Stdout
		} else {
			out, err = os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("Failed to open log file: %v", err)
			}
		}

		instance = &Logger{
			logger: log.New(out, "", log.LstdFlags|log.Lshortfile),
			level:  level,
		}
	})
	return instance, err
}

// log formats and writes log messages
func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := logLevels[level]
	color := logColors[level]
	message := fmt.Sprintf(format, v...)

	logLine := fmt.Sprintf("%s[%s] [%s] %s%s", color, timestamp, levelStr, message, ColorReset)

	l.logger.Println(logLine)
}

// Logging methods
func (l *Logger) Debug(format string, v ...interface{}) { l.log(DebugLevel, format, v...) }
func (l *Logger) Info(format string, v ...interface{})  { l.log(InfoLevel, format, v...) }
func (l *Logger) Warn(format string, v ...interface{})  { l.log(WarnLevel, format, v...) }
func (l *Logger) Error(format string, v ...interface{}) { l.log(ErrorLevel, format, v...) }
func (l *Logger) SetOutput(out io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.SetOutput(out)
}
// Global logging methods for convenience
func Debug(format string, v ...interface{}) { GetLogger().Debug(format, v...) }
func Info(format string, v ...interface{})  { GetLogger().Info(format, v...) }
func Warn(format string, v ...interface{})  { GetLogger().Warn(format, v...) }
func Error(format string, v ...interface{}) { GetLogger().Error(format, v...) }

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	if instance == nil {
		log.Fatalf("Logger not initialized. Call InitLogger first.")
	}
	return instance
}
