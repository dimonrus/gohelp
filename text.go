package gohelp

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var matchUUIDPattern = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[rand.Int63()%int64(len(Charset))]
	}
	return string(b)
}

type UUID struct {
	value *string
}

func NewUUID() string {
	return *(&UUID{}).Generate().Get()
}

func (u *UUID) Generate() *UUID {
	b := make([]byte, 16)
	_, err := crand.Read(b)
	if err != nil {
		return nil
	}
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	u.value = &uuid
	return u
}

func (u *UUID) Reset() *UUID {
	u.value = nil
	return u
}

func (u *UUID) Get() *string {
	return u.value
}

func (u *UUID) IsValid() bool {
	return matchUUIDPattern.MatchString(*u.Get())
}

func ParseUUID(data string) *UUID {
	uuid := &UUID{value: &data}
	if uuid.IsValid() {
		return uuid
	}
	return nil
}

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
		}
	}
	return buf.String()
}

func BeforeString(source string, substr string) string {
	// Get substring before a string.
	pos := strings.Index(source, substr)
	if pos == -1 {
		return ""
	}
	return source[0:pos]
}
