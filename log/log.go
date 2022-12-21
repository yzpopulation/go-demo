package log

import (
	"fmt"
	"os"
	"time"
)

const colorBlack = "\033[0;30m"
const colorBlackL = "\033[1;30m"
const colorRed = "\033[0;31m"
const colorRedL = "\033[1;31m"
const colorGreen = "\033[0;32m"
const colorGreenL = "\033[1;32m"
const colorBrown = "\033[0;33m"
const colorYellow = "\033[1;33m"
const colorBlue = "\033[0;34m"
const colorBlueL = "\033[1;34m"
const colorPurple = "\033[0;35m"
const colorPurpleL = "\033[1;35m"
const colorCyan = "\033[0;36m"
const colorCyanL = "\033[1;36m"
const colorWhite = "\033[0;37m"
const colorWhiteL = "\033[1;37m"
const colorNone = "\033[0m"

func Debug(log string) {
	fmt.Fprintf(os.Stdout, "%s[DBG][%s] %s%s\n", colorBlack, getTime(), log, colorNone)
}
func Info(log string) {
	fmt.Fprintf(os.Stdout, "%s[INF][%s] %s%s\n", colorWhite, getTime(), log, colorNone)
}
func Warn(log string) {
	fmt.Fprintf(os.Stdout, "%s[WAR][%s] %s%s\n", colorBrown, getTime(), log, colorNone)
}
func Error(log string) {
	fmt.Fprintf(os.Stdout, "%s[ERR][%s] %s%s\n", colorRed, getTime(), log, colorNone)
}
func Fatal(log string) {
	fmt.Fprintf(os.Stdout, "%s[FAT][%s] %s%s\n", colorPurple, getTime(), log, colorNone)
}
func Panic(log string) {
	fmt.Fprintf(os.Stdout, "%s[PNC][%s] %s%s\n", colorPurple, getTime(), log, colorNone)
}

func Default(log string) {
	Info(log)
}

func getTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}
