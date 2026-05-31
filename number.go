package gohelp

import (
	"crypto/rand"
	"math/big"
	"time"
)

func GetRndDateTime() time.Time {
	minimum := int(time.Date(2001, 1, 0, 0, 0, 0, 0, time.UTC).Unix())
	maximum := int(time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix())
	return time.Unix(int64(GetRndNumber(minimum, maximum)), 0)
}

func GetRndId() int {
	minimum := 77770000000
	maximum := 87770000000
	return GetRndNumber(minimum, maximum)
}

func GetRndPhone() int {
	minimum := 1000000000
	maximum := 2000000000
	return GetRndNumber(minimum, maximum)
}

func GetRndNumber(min int, max int) int {
	rangeLimit := big.NewInt(int64(max - min + 1))
	n, _ := rand.Int(rand.Reader, rangeLimit)
	result := n.Int64() + int64(min)
	return int(result)
}
