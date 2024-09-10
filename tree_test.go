package gohelp

import (
	"encoding/json"
	"testing"
)

type TestTreeData struct {
	Name     string `json:"name" valid:"required"`
	Color    int32  `json:"color" valid:"required"`
	IsHidden bool   `json:"isHidden"`
}

func getTestTree() Tree[TestTreeData] {
	treeListData := make(TreeList[TestTreeData], 4)

	t1 := Tree[TestTreeData]{
		Id: 1,
		Data: &TestTreeData{
			Name:     "foo",
			Color:    112231,
			IsHidden: false,
		},
		Children: treeListData[:2],
	}
	t2 := Tree[TestTreeData]{
		Id:     2,
		Parent: &t1,
		Data: &TestTreeData{
			Name:     "bar",
			Color:    112232,
			IsHidden: true,
		},
	}
	t3 := Tree[TestTreeData]{
		Id:     3,
		Parent: &t1,
		Data: &TestTreeData{
			Name:     "baz",
			Color:    112233,
			IsHidden: false,
		},
		Children: treeListData[2:3],
	}
	t4 := Tree[TestTreeData]{
		Id:     4,
		Parent: &t3,
		Data: &TestTreeData{
			Name:     "baz3",
			Color:    112234,
			IsHidden: false,
		},
		Children: treeListData[3:4],
	}
	t5 := Tree[TestTreeData]{
		Id:     5,
		Parent: &t4,
		Data: &TestTreeData{
			Name:     "baz4",
			Color:    112235,
			IsHidden: true,
		},
	}
	treeListData[0] = &t2
	treeListData[1] = &t3
	treeListData[2] = &t4
	treeListData[3] = &t5
	return t1
}

func TestTree(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		tree := getTestTree()
		if tree.FirstChild().Id != 2 {
			t.Fatal("wrong")
		}
	})
	t.Run("last", func(t *testing.T) {
		tree := getTestTree()
		if tree.LastChild().Id != 3 {
			t.Fatal("wrong")
		}
	})
	t.Run("list.GetById", func(t *testing.T) {
		tree := getTestTree()
		list := TreeList[TestTreeData]{&tree}
		item := list.GetById(5)
		if item == nil || item.Id != 5 {
			t.Fatal("wrong")
		}
	})
	t.Run("item.GetById", func(t *testing.T) {
		tree := getTestTree()
		item := tree.GetById(5)
		if item == nil || item.Id != 5 {
			t.Fatal("wrong")
		}
	})
	t.Run("path", func(t *testing.T) {
		tree := getTestTree()
		item := tree.GetById(5)
		if item == nil || item.Id != 5 {
			t.Fatal("wrong")
		}
		path := item.Path(" / ", func(i *Tree[TestTreeData]) string {
			return i.Data.Name
		})
		if path != "foo / baz / baz3 / baz4" {
			t.Fatal("wrong path")
		}
	})
	t.Run("lookup.list", func(t *testing.T) {
		tree := getTestTree()
		list := TreeList[TestTreeData]{&tree}
		items := list.Lookup(func(i *Tree[TestTreeData]) bool {
			if i.Id == 3 || i.Data.Name == "bar" {
				return true
			}
			return false
		})
		if len(items) != 2 {
			t.Fatal("wrong item count")
		}
		var foundId3, foundBaz bool
		for _, item := range items {
			if item.Id == 3 {
				foundId3 = true
			}
			if item.Data.Name == "bar" {
				foundBaz = true
			}
		}
		if !foundId3 || !foundBaz {
			t.Fatal("wrong lookup")
		}
	})
	t.Run("lookup.item", func(t *testing.T) {
		tree := getTestTree()
		items := tree.Lookup(func(i *Tree[TestTreeData]) bool {
			if i.Id == 3 || i.Data.Name == "baz3" {
				return true
			}
			return false
		})
		if len(items) != 2 {
			t.Fatal("wrong item count")
		}
		var foundId3, foundBaz bool
		for _, item := range items {
			if item.Id == 3 {
				foundId3 = true
			}
			if item.Data.Name == "baz3" {
				foundBaz = true
			}
		}
		if !foundId3 || !foundBaz {
			t.Fatal("wrong lookup")
		}
	})
	t.Run("json.marshal", func(t *testing.T) {
		tree := getTestTree()
		data, err := json.Marshal(tree)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(data))
	})
	t.Run("item.map", func(t *testing.T) {
		tree := getTestTree()
		var count int
		tree.Map(func(i *Tree[TestTreeData]) {
			count++
			i.Data.Name = "__" + i.Data.Name + "__"
		})
		node := tree.GetById(5)
		path := node.Path(" / ", func(i *Tree[TestTreeData]) string {
			return i.Data.Name
		})
		if path != "__foo__ / __baz__ / __baz3__ / __baz4__" {
			t.Fatal("wrong map")
		}
		t.Log(count)
	})
	t.Run("item.unmarshal", func(t *testing.T) {
		tr := getTestTree()
		data, err := json.Marshal(tr)
		tree := &Tree[TestTreeData]{}
		err = tree.UnmarshalJSON(data)
		if err != nil {
			t.Fatal(err)
		}
	})

}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gohelp
// cpu: Apple M2 Max
// BenchmarkTree_Map
// BenchmarkTree_Map-12    	 3439864	       336.5 ns/op	     472 B/op	      17 allocs/op
func BenchmarkTree_Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := getTestTree()
		_ = tree
		tree.Map(func(i *Tree[TestTreeData]) {
			i.Data.Name += "__"
		})
	}
	b.ReportAllocs()
}
