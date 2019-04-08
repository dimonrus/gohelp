package gohelp

import (
	"math/rand"
	"time"
)

func GetRndDateTime() time.Time {
	min := time.Date(2001, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	rand.Seed(time.Now().UTC().UnixNano())
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GetRndId() int {
	min := 77770000000
	max := 87770000000
	return GetRndNumber(min, max)
}

func GetRndPhone() int {
	min := 1000000000
	max := 2000000000
	return GetRndNumber(min, max)
}

func GetRndNumber(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
