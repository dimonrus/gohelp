package gohelp

import (
	"bytes"
	"encoding/gob"
	"sync"
)

// PipeInstruction Instruction declaration
type PipeInstruction []byte

// PipeArguments list of Arguments
type PipeArguments []any

// Pipe single call
type Pipe struct {
	// Instruction for call
	Instruction PipeInstruction
	// argument for call
	Arguments []PipeArguments
}

// Reset reset pipe
func (d *Pipe) Reset() {
	d.Instruction = d.Instruction[0:]
	d.Arguments = d.Arguments[0:]
	return
}

// Push Instruction to pipeline with arguments
func (d *Pipe) Push(instruction PipeInstruction, args ...any) {
	d.Instruction = instruction
	d.Arguments = append(d.Arguments, args)
	return
}

// NewPipe Init pipe
func NewPipe(lenI, lenA int) Pipe {
	return Pipe{
		Instruction: make(PipeInstruction, 0, lenI),
		Arguments:   make([]PipeArguments, 0, lenA),
	}
}

// PipeLine collected pipe items
type PipeLine struct {
	// length of instructions in bytes
	lenI int
	// length of Arguments list
	lenA int
	// all pipe items
	items []Pipe
	// mutex
	m sync.RWMutex
}

// Serialize pipeline
func (s *PipeLine) Serialize(buf *bytes.Buffer) ([]byte, error) {
	g := gob.NewEncoder(buf)
	s.m.Lock()
	err := g.Encode(s.items)
	s.m.Unlock()
	return buf.Bytes(), err
}

// Deserialize pipeline
func (s *PipeLine) Deserialize(buf *bytes.Buffer) error {
	g := gob.NewDecoder(buf)
	return g.Decode(&s.items)
}

// Add save new call
func (s *PipeLine) Add(instruction PipeInstruction, args ...any) {
	call := NewPipe(s.lenI, s.lenA)
	call.Push(instruction, args...)
	s.m.Lock()
	s.items = append(s.items, call)
	s.m.Unlock()
}

// Current save into current call
func (s *PipeLine) Current(instruction PipeInstruction, args ...any) {
	s.m.Lock()
	(s.items)[len(s.items)-1].Push(instruction, args...)
	s.m.Unlock()
}

// Reset reset
func (s *PipeLine) Reset() {
	s.items = s.items[:0]
	return
}

// Items item iterator
func (s *PipeLine) Items() []Pipe {
	return s.items
}

// NewPipeLine Create new pipeline with arguments
// lenI len of instruction field
// lenA len of arguments list
func NewPipeLine(lenI, lenA int) *PipeLine {
	return &PipeLine{lenI: lenI, lenA: lenA, items: make([]Pipe, 0, 10)}
}
