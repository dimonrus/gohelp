package gohelp

import (
	"strconv"
	"strings"
)

// uint8
// uint16
// uint32
// uint64
// int8
// int16
// int32
// int64
// float32
// float64
// string
// int
// uint
// uintptr
// byte
// rune

func ExistsInArrayString(value string, slice []string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}

func ExistsInArrayInt(value int, slice []int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}

func ExistsInArrayInt32(value int32, slice []int32) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}

func ExistsInArrayInt64(value int64, slice []int64) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}

func ArrayIntersectString(target []string, destination []string) (intersect []string, removedTargets int) {
	for _, tItem := range target {
		targetFound := false
		for _, dItem := range destination {
			if tItem == dItem {
				intersect = append(intersect, tItem)
				targetFound = true
				break
			}
		}
		if targetFound == false {
			removedTargets += 1
		}
	}

	return intersect, removedTargets
}

func ImplodeUint64(values []uint64, sep string) string {
	if len(values) == 0 {
		return ""
	}

	b := make([]string, len(values))
	for i, v := range values {
		b[i] = strconv.FormatUint(v, 10)
	}

	return strings.Join(b, sep)
}

func ImplodeInt64(values []int64, sep string) string {
	if len(values) == 0 {
		return ""
	}

	b := make([]string, len(values))
	for i, v := range values {
		b[i] = strconv.FormatInt(v, 10)
	}

	return strings.Join(b, sep)
}

func ImplodeInt(values []int, sep string) string {
	if len(values) == 0 {
		return ""
	}

	b := make([]string, len(values))
	for i, v := range values {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, sep)
}

func Implode(values []interface{}, sep string) string {
	if len(values) == 0 {
		return ""
	}

	b := make([]string, len(values))
	for i, v := range values {
		switch values[i].(type) {
		case uint64:
			b[i] = strconv.FormatUint(v.(uint64), 10)
		case uint32:
			b[i] = strconv.FormatUint(uint64(v.(uint32)), 10)
		case uint16:
			b[i] = strconv.FormatUint(uint64(v.(uint16)), 10)
		case uint8:
			b[i] = strconv.FormatUint(uint64(v.(uint8)), 10)
		case uint:
			b[i] = strconv.FormatUint(uint64(v.(uint)), 10)
		case int64:
			b[i] = strconv.FormatInt(v.(int64), 10)
		case int32:
			b[i] = strconv.FormatInt(int64(v.(int32)), 10)
		case int16:
			b[i] = strconv.FormatInt(int64(v.(int16)), 10)
		case int8:
			b[i] = strconv.FormatInt(int64(v.(int8)), 10)
		case int:
			b[i] = strconv.FormatInt(int64(v.(int)), 10)
		}
	}

	return strings.Join(b, sep)
}

func AppendUniqueUint8(slice []uint8, values ...uint8) []uint8 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueUint16(slice []uint16, values ...uint16) []uint16 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueUint32(slice []uint32, values ...uint32) []uint32 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueUint64(slice []uint64, values ...uint64) []uint64 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueInt8(slice []int8, values ...int8) []int8 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueInt16(slice []int16, values ...int16) []int16 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueInt32(slice []int32, values ...int32) []int32 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueInt64(slice []int64, values ...int64) []int64 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueFloat32(slice []float32, values ...float32) []float32 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueFloat64(slice []float64, values ...float64) []float64 {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueString(slice []string, values ...string) []string {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueInt(slice []int, values ...int) []int {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueUint(slice []uint, values ...uint) []uint {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueUintptr(slice []uintptr, values ...uintptr) []uintptr {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueByte(slice []byte, values ...byte) []byte {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}
func AppendUniqueRune(slice []rune, values ...rune) []rune {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}

func UniqueLeftUint8(left []uint8, right []uint8) []uint8 {
	var leftUnique = make([]uint8, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueUint8([]uint8{}, leftUnique...)
}
func UniqueLeftUint16(left []uint16, right []uint16) []uint16 {
	var leftUnique = make([]uint16, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueUint16([]uint16{}, leftUnique...)
}
func UniqueLeftUint32(left []uint32, right []uint32) []uint32 {
	var leftUnique = make([]uint32, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueUint32([]uint32{}, leftUnique...)
}
func UniqueLeftUint64(left []uint64, right []uint64) []uint64 {
	var leftUnique = make([]uint64, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueUint64([]uint64{}, leftUnique...)
}
func UniqueLeftInt8(left []int8, right []int8) []int8 {
	var leftUnique = make([]int8, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueInt8([]int8{}, leftUnique...)
}
func UniqueLeftInt16(left []int16, right []int16) []int16 {
	var leftUnique = make([]int16, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueInt16([]int16{}, leftUnique...)
}
func UniqueLeftInt32(left []int32, right []int32) []int32 {
	var leftUnique = make([]int32, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueInt32([]int32{}, leftUnique...)
}
func UniqueLeftInt64(left []int64, right []int64) []int64 {
	var leftUnique = make([]int64, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueInt64([]int64{}, leftUnique...)
}
func UniqueLeftFloat32(left []float32, right []float32) []float32 {
	var leftUnique = make([]float32, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueFloat32([]float32{}, leftUnique...)
}
func UniqueLeftFloat64(left []float64, right []float64) []float64 {
	var leftUnique = make([]float64, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueFloat64([]float64{}, leftUnique...)
}
func UniqueLeftString(left []string, right []string) []string {
	var leftUnique = make([]string, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueString([]string{}, leftUnique...)
}
func UniqueLeftInt(left []int, right []int) []int {
	var leftUnique = make([]int, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueInt([]int{}, leftUnique...)
}
func UniqueLeftUint(left []uint, right []uint) []uint {
	var leftUnique = make([]uint, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueUint([]uint{}, leftUnique...)
}
func UniqueLeftUintptr(left []uintptr, right []uintptr) []uintptr {
	var leftUnique = make([]uintptr, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueUintptr([]uintptr{}, leftUnique...)
}
func UniqueLeftByte(left []byte, right []byte) []byte {
	var leftUnique = make([]byte, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueByte([]byte{}, leftUnique...)
}
func UniqueLeftRune(left []rune, right []rune) []rune {
	var leftUnique = make([]rune, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUniqueRune([]rune{}, leftUnique...)
}

type Int32Slice []int32

func (x Int32Slice) Len() int           { return len(x) }
func (x Int32Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Int32Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Int64Slice []int32

func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
