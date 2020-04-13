package gohelp

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchUUIDPattern = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

func RandString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[source.Int63()%int64(len(Charset))]
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

func ToUnderscore(str string) string {
	underscore := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	underscore = matchAllCap.ReplaceAllString(underscore, "${1}_${2}")
	return strings.ToLower(underscore)
}

func ToCamelCase(str string, isFirstTitle bool) (string, error) {
	var result string
	reg, _ := regexp.Compile(`[a-z0-9]+(_[a-z0-9]+)*`)
	regs, _ := regexp.Compile(`\s*`)
	matches := reg.FindStringSubmatch(strings.ToLower(str))

	if len(matches) > 0 {
		keys := strings.Split(matches[0], "_")
		var titled string
		if isFirstTitle {
			titled = strings.Join(keys, " ")
		} else {
			titled = strings.Join(keys[1:], " ")
		}
		str := reg.ReplaceAllStringFunc(titled, func(s string) string {
			return strings.Title(s)
		})
		str = regs.ReplaceAllString(str, "")
		if !isFirstTitle {
			result = keys[0] + str
		} else {
			result = str
		}
	} else {
		return "", errors.New("wrong string passed")
	}
	return result, nil
}

func BeforeString(source string, substr string) string {
	// Get substring before a string.
	pos := strings.Index(source, substr)
	if pos == -1 {
		return ""
	}
	return source[0:pos]
}
