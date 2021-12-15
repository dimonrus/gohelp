package gohelp

// ByteStack Stack of bytes
type ByteStack [][]byte

// Pop value
func (s ByteStack) Pop() ([]byte, ByteStack) {
	if len(s) > 0 {
		return s[len(s)-1], s[:len(s)-1]
	}
	return nil, s
}

// Push value
func (s ByteStack) Push(v []byte) ByteStack {
	return append(s, v)
}

