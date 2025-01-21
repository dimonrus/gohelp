package gohelp

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPipeLine_Serialize(t *testing.T) {
	t.Run("serial", func(t *testing.T) {
		ds := NewPipeLine(100, 1000)
		ds.Add(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 10, "some", false)
		ds.Current(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 15, "sometwo", true)
		ds.Add(PipeInstruction("DELETE FROM organisation WHERE ids = ?"), []int{10, 2, 3})
		buf := bytes.NewBuffer([]byte{})
		data, err := ds.Serialize(buf)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%X", data)
		ds.Reset()
		out := NewPipeLine(100, 1000)
		desb := bytes.NewBuffer(data)
		err = out.Deserialize(desb)
		if err != nil {
			t.Fatal(err)
		}
		if string(out.items[0].Instruction) != "SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?" {
			t.Fatal("Instruction is not saved")
		}
		if (out.items[0].Arguments[0][0]).(int) != 10 {
			t.Fatal("Instruction is not saved")
		}
		if (out.items[1].Arguments[0][0]).([]int)[0] != 10 {
			t.Fatal("arguments array is not saved properly")
		}
		for _, item := range out.Items() {
			t.Log(string(item.Instruction), len(item.Arguments))
			item.Reset()
		}
	})
}

func BenchmarkPipeLine(b *testing.B) {
	b.Run("serialize", func(b *testing.B) {
		ds := NewPipeLine(100, 1000)
		ds.Add(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 10, "some", false)
		ds.Current(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 15, "sometwo", true)
		ds.Add(PipeInstruction("DELETE FROM organisation WHERE ids = ?"), []int{10, 2, 3})
		buf := bytes.NewBuffer([]byte{})
		for i := 0; i < b.N; i++ {
			ds.Serialize(buf)
			ds.Reset()
		}
		b.ReportAllocs()
	})
	b.Run("deserialize", func(b *testing.B) {
		ds := NewPipeLine(100, 1000)
		ds.Add(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 10, "some", false)
		ds.Current(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 15, "sometwo", true)
		ds.Add(PipeInstruction("DELETE FROM organisation WHERE ids = ?"), []int{10, 2, 3})
		out := NewPipeLine(100, 1000)
		buf := bytes.NewBuffer([]byte{})
		data, _ := ds.Serialize(buf)
		des := bytes.NewBuffer(data)
		for i := 0; i < b.N; i++ {
			out.Deserialize(des)
			out.Reset()
		}
		b.ReportAllocs()
	})
	b.Run("lot_of_arguments", func(b *testing.B) {
		ds := NewPipeLine(100, 1000)
		ds.Add(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), 10, "some", false)
		for i := 0; i < 1000000; i++ {
			ds.Current(PipeInstruction("SELECT * FROM person WHERE id = ? AND name = ? AND is_deleted = ?"), i, "some_"+fmt.Sprintf("%v", i), i%4 == 0)
		}
		buf := bytes.NewBuffer([]byte{})
		for i := 0; i < b.N; i++ {
			ds.Serialize(buf)
			ds.Reset()
		}
		b.ReportAllocs()
	})
	b.Run("lot_of_instructions", func(b *testing.B) {
		ds := NewPipeLine(100, 1000)
		for i := 0; i < 10000; i++ {
			ds.Add(PipeInstruction("CREATE TABLE IF NOT EXIST some_table (id int, name TEXT, data JSON)"))
		}
		buf := bytes.NewBuffer([]byte{})
		for i := 0; i < b.N; i++ {
			ds.Serialize(buf)
			ds.Reset()
		}
		b.ReportAllocs()
	})
}
