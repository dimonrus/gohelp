package gohelp

import (
	"strconv"
	"strings"
)

func AppendUniqueUint64(slice []uint64, i uint64) []uint64 {
	for _, el := range slice {
		if el == i {
			return slice
		}
	}
	return append(slice, i)
}

func ExistsInArrayString(value string, slice []string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}

func AppendUniqueInt(slice []int, i int) []int {
	for _, el := range slice {
		if el == i {
			return slice
		}
	}
	return append(slice, i)
}

func AppendUniqueInt64(slice []int64, i int64) []int64 {
	for _, el := range slice {
		if el == i {
			return slice
		}
	}
	return append(slice, i)
}

func AppendUniqueString(slice []string, i string) []string {
	for _, el := range slice {
		if el == i {
			return slice
		}
	}
	return append(slice, i)
}

func ExistsInArrayInt(value int, slice []int) bool {
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
