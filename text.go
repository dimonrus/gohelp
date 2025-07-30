package gohelp

import (
	"errors"
	"strings"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString create random string
func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[rnd.Int63()%int64(len(Charset))]
	}
	return string(b)
}

// ToUnderscore transform string to underscore case
func ToUnderscore(str string) string {
	var buf = make([]byte, 0, len(str)+8)
	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			if len(buf) > 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(c-'A'+'a'))
		} else {
			buf = append(buf, byte(c))
		}
	}
	return Convert[[]byte, string](buf)
}

// ToCamelCase transform to camelCase
func ToCamelCase(str string, isFirstTitle bool) string {
	var buf = make([]byte, 0, len(str))
	for i, c := range str {
		if 'a' <= c && c <= 'z' {
			if len(buf) == 0 {
				if isFirstTitle {
					buf = append(buf, byte(c-'a'+'A'))
				} else {
					buf = append(buf, byte(c))
				}
			} else {
				if str[i-1] == '_' {
					buf = append(buf, byte(c-'a'+'A'))
				} else {
					buf = append(buf, byte(c))
				}
			}
		} else if 'A' <= c && c <= 'Z' {
			if len(buf) == 0 {
				if isFirstTitle {
					buf = append(buf, byte(c))
				} else {
					buf = append(buf, byte(c+'a'-'A'))
				}
			} else {
				buf = append(buf, byte(c))
			}
		} else if '0' <= c && c <= '9' {
			buf = append(buf, byte(c))
		}
	}
	return Convert[[]byte, string](buf)
}

// BeforeString get string in source before substring
func BeforeString(source string, substr string) string {
	// Get substring before a string.
	pos := strings.Index(source, substr)
	if pos == -1 {
		return ""
	}
	return source[0:pos]
}

// CheckBracers check if bracers are correct
func CheckBracers(source string, stack Stack[byte]) error {
	for _, s := range source {
		switch s {
		case '{', '[', '(':
			stack = stack.Push(byte(s))
		case ']', '}', ')':
			var exists bool
			var bracer byte
			bracer, exists, stack = stack.Pop()
			if !exists {
				return errors.New("incorrect closed bracers count")
			}
			if bracer == '{' && s != '}' {
				return errors.New("incorrect figure bracer found")
			}
			if bracer == '[' && s != ']' {
				return errors.New("incorrect square pair bracer found")
			}
			if bracer == '(' && s != ')' {
				return errors.New("incorrect circle pair bracer found")
			}
		}
	}
	if len(stack) != 0 {
		return errors.New("incorrect bracers count")
	}
	return nil
}

// CheckTypeOf check value for type
// 00001111
// - 1 - uint
// - 2 - int
// - 4 - float
// - 8 - string
func CheckTypeOf(value []byte) (isUint, isInt, isFloat, isBool, isString bool) {
	var valueType uint8
	for _, b := range value {
		if b >= '0' && b <= '9' {
			valueType = valueType | 1
		} else if b == '-' {
			if valueType >= 1 {
				valueType ^= 8
			}
			valueType = valueType | 2
		} else if b == '.' {
			if valueType >= 4 {
				valueType ^= 8
			}
			valueType = valueType | 4
		} else {
			valueType = valueType | 8
		}
	}
	if valueType&8 == 8 {
		lower := strings.ToLower(string(value))
		if lower == "true" || lower == "false" {
			isBool = true
			return
		}
		isString = true
		return
	}
	if valueType&4 == 4 {
		isFloat = true
		return
	}
	if valueType&2 == 2 {
		isInt = true
		return
	}
	if valueType&1 == 1 {
		isUint = true
		return
	}
	if valueType == 0 {
		isString = true
	}
	return
}
