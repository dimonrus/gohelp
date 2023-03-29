package gohelp

import (
	"errors"
	"math/rand"
	"strings"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString create random string
func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[rand.Int63()%int64(len(Charset))]
	}
	return string(b)
}

// ToUnderscore transform string to underscore case
func ToUnderscore(str string) string {
	var buf strings.Builder
	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c - 'A' + 'a')
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

// ToCamelCase transform to camelCase
func ToCamelCase(str string, isFirstTitle bool) string {
	var buf strings.Builder
	for i, c := range str {
		if 'a' <= c && c <= 'z' {
			if buf.Len() == 0 {
				if isFirstTitle {
					buf.WriteRune(c - 'a' + 'A')
				} else {
					buf.WriteRune(c)
				}
			} else {
				if str[i-1] == '_' {
					buf.WriteRune(c - 'a' + 'A')
				} else {
					buf.WriteRune(c)
				}
			}
		} else if 'A' <= c && c <= 'Z' {
			if buf.Len() == 0 {
				if isFirstTitle {
					buf.WriteRune(c)
				} else {
					buf.WriteRune(c + 'a' - 'A')
				}
			} else {
				buf.WriteRune(c)
			}
		}
	}
	return buf.String()
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

// CheckBracers check if bracers is corrects
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
