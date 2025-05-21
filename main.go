package logger

import (
	"fmt"
	"os"
	"time"
)

var DisableWarn = false
var DisableInfo = false
var DisableDebug = false
var DisableSuccess = false
var EnableTimeStamp = false

func formatMessage(msg string) string {
	if EnableTimeStamp {
		timeStamp := time.Now().Format("2006-01-02 15:04:05")
		return fmt.Sprintf("[%s] %s", timeStamp, msg)
	}
	return msg
}

func Error(message string) {
	fmt.Printf("%s[ERROR]%s %s\n", Red, Reset, formatMessage(message))
	os.Exit(1)
}

func Errorf(format string, a ...any) {
	Error(fmt.Sprintf(format, a...))
}

func Info(message string) {
	if DisableInfo {
		return
	}
	fmt.Printf("%s[INFO]%s %s\n", LightBlue, Reset, formatMessage(message))
}

func Infof(format string, a ...any) {
	Info(fmt.Sprintf(format, a...))
}

func Success(message string) {
	if DisableSuccess {
		return
	}
	fmt.Printf("%s[OK]%s %s\n", LightGreen, Reset, formatMessage(message))
}

func Successf(format string, a ...any) {
	Success(fmt.Sprintf(format, a...))
}

func Warn(message string) {
	if DisableWarn {
		return
	}
	fmt.Printf("%s[WARNING]%s %s\n", Yellow, Reset, formatMessage(message))
}

func Warnf(format string, a ...any) {
	Warn(fmt.Sprintf(format, a...))
}

func Debug(message string) {
	if DisableDebug {
		return
	}
	fmt.Printf("%s[DEBUG]%s %s\n", Cyan, Reset, formatMessage(message))
}

func Debugf(format string, a ...any) {
	Debug(fmt.Sprintf(format, a...))
}
