package gohelp

import "fmt"

const (
	AnsiReset = "\x1b[0m"

	AnsiBlack   = "\x1b[30;1m"
	AnsiRed     = "\x1b[31;1m"
	AnsiGreen   = "\x1b[32;1m"
	AnsiYellow  = "\x1b[33;1m"
	AnsiBlue    = "\x1b[34;1m"
	AnsiMagenta = "\x1b[35;1m"
	AnsiCyan    = "\x1b[36;1m"
	AnsiWhite   = "\x1b[37;1m"

	AnsiCustom = "\x1b[38;5;%vm"

	AnsiBackgroundBlack   = "\x1b[40;1m"
	AnsiBackgroundRed     = "\x1b[42;1m"
	AnsiBackgroundGreen   = "\x1b[42;1m"
	AnsiBackgroundYellow  = "\x1b[43;1m"
	AnsiBackgroundBlue    = "\x1b[44;1m"
	AnsiBackgroundMagenta = "\x1b[45;1m"
	AnsiBackgroundCyan    = "\x1b[46;1m"
	AnsiBackgroundWhite   = "\x1b[47;1m"

	AnsiBackgroundCustom = "\x1b[48;5;%vm"

	AnsiBold      = "\x1b[1m"
	AnsiUnderline = "\x1b[4m"
	AnsiReversed  = "\x1b[7m"

	AnsiCursorUp    = "\x1b[%vA"
	AnsiCursorDown  = "\x1b[%vB"
	AnsiCursorRight = "\x1b[%vC"
	AnsiCursorLeft  = "\x1b[%vD"
)

// Red wrap into red color
func Red(v interface{}) string {
	return AnsiRed + fmt.Sprintf("%v", v) + AnsiReset
}

// Yellow wrap into yellow color
func Yellow(v interface{}) string {
	return AnsiYellow + fmt.Sprintf("%v", v) + AnsiReset
}

// Green wrap into green color
func Green(v interface{}) string {
	return AnsiGreen + fmt.Sprintf("%v", v) + AnsiReset
}

// Blue wrap into blue color
func Blue(v interface{}) string {
	return AnsiBlue + fmt.Sprintf("%v", v) + AnsiReset
}

// Magenta wrap into magenta color
func Magenta(v interface{}) string {
	return AnsiMagenta + fmt.Sprintf("%v", v) + AnsiReset
}

// Cyan wrap into cyan color
func Cyan(v interface{}) string {
	return AnsiCyan + fmt.Sprintf("%v", v) + AnsiReset
}

// Underline wrap into underline text
func Underline(v interface{}) string {
	return AnsiUnderline + fmt.Sprintf("%v", v) + AnsiReset
}
