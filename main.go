package logger

import (
	"fmt"
)

func LogError(message string) {
	fmt.Printf("%s[ERROR]%s %s\n", Red, Reset, message)
}

func LogInfo(message string) {
	fmt.Printf("%s[INFO]%s %s\n", LightBlue, Reset, message)
}

func LogSuccess(message string) {
	fmt.Printf("%s[OK]%s %s\n", LightGreen, Reset, message)
}

func LogWarning(message string) {
	fmt.Printf("%s[WARNING]%s %s\n", Yellow, Reset, message)
}

func LogDebug(message string) {
	fmt.Printf("%s[DEBUG]%s %s\n", Cyan, Reset, message)
}
