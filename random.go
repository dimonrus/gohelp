package gohelp

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

// RandInt64 get random int64
func RandInt64() (int64, error) {
	u, err := RandUnt64()
	if err != nil {
		return 0, err
	}
	if u >= (1<<63 - 1) {
		return -int64(1<<63 - 1), nil
	}
	return int64(u), nil
}

// RandInt63n Get none negative random int64
func RandInt63n() (int64, error) {
	rangeLimit := big.NewInt(int64(1<<63 - 1))
	n, _ := rand.Int(rand.Reader, rangeLimit)
	return n.Int64(), nil
}

// RandUnt64 get random uint64
func RandUnt64() (uint64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}
