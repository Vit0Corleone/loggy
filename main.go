package loggy

import (
	"fmt"
	"strings"
	"time"

	"github.com/Vit0Corleone/loggy/colors"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"

	TextFormat = "text"
	JSONFormat = "json"
)

var Time = colors.SetColor(time.Now().Format("2006-01-02 15:04:05.00"), colors.DarkBlue)

type Logger struct {
	Level  string
	Format string
}

func NewLogger(level, format string) *Logger {
	log := &Logger{
		Level:  level,
		Format: format,
	}

	return log
}

func (l *Logger) Debug(msg string) {
	if l.Level != LevelDebug {
		return
	}
	level := colors.SetColor(strings.ToUpper(LevelDebug), colors.White)
	msg = colors.SetColor(msg, colors.White)
	logIt(level, l.Format, msg)
}

func (l *Logger) Info(msg string) {
	if l.Level == LevelWarn || l.Level == LevelError {
		return
	}
	level := colors.SetColor(strings.ToUpper(LevelInfo), colors.Green)
	msg = colors.SetColor(msg, colors.Cyan)
	logIt(level, l.Format, msg)
}

func (l *Logger) Warn(msg string) {
	if l.Level == LevelError {
		return
	}
	level := colors.SetColor(strings.ToUpper(LevelWarn), colors.Orange)
	msg = colors.SetColor(msg, colors.Orange)
	logIt(level, l.Format, msg)
}

func (l *Logger) Error(msg string) {
	level := colors.SetColor(strings.ToUpper(LevelError), colors.Red)
	msg = colors.SetColor(msg, colors.Red)
	logIt(level, l.Format, msg)
}

func logIt(level, format, msg string) {
	switch format {
	case TextFormat:
		fmt.Printf("%v\n[%s]: %s\n\n", Time, level, msg)
	default:
		fmt.Printf("{\n    TIME: %v,\n    LEVEL: %s,\n    MESSAGE: %s,\n},\n", Time, level, msg)
	}
}
