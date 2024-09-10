package gohelp

import "encoding/json"

// TreeList list of tree items
type TreeList[T any] []*Tree[T]

// GetById find by id
func (l TreeList[T]) GetById(id int) *Tree[T] {
	for i := range l {
		if l[i].Id == id {
			return l[i]
		}
		if len(l[i].Children) > 0 {
			return l[i].Children.GetById(id)
		}
	}
	return nil
}

// Lookup function
func (l TreeList[T]) Lookup(found func(i *Tree[T]) bool) TreeList[T] {
	var result = make(TreeList[T], 0, 4)
	for i := range l {
		if found(l[i]) {
			result = append(result, l[i])
		}
		if len(l[i].Children) > 0 {
			result = append(result, l[i].Children.Lookup(found)...)
		}
	}
	return result
}

// First tree item
func (l TreeList[T]) First() *Tree[T] {
	if len(l) == 0 {
		return nil
	}
	return l[0]
}

// Last tree item
func (l TreeList[T]) Last() *Tree[T] {
	if len(l) == 0 {
		return nil
	}
	return l[len(l)-1]
}

// Tree simple tree
type Tree[T any] struct {
	// Identifier
	Id int `json:"id"`
	// Parent item
	Parent *Tree[T] `json:"-"`
	// Specific tree data
	Data *T `json:"data"`
	// Children
	Children TreeList[T] `json:"children"`
}

// Tree for json unmarshal
type jsonTree[T any] struct {
	// Identifier
	Id int `json:"id"`
	// Specific tree data
	Data *T `json:"data"`
	// Children
	Children jsonTreeList[T] `json:"children"`
}

// jsonTreeList for unmarshal
type jsonTreeList[T any] []*jsonTree[T]

// UnmarshalJSON unmarshal data with parents
func (t *Tree[T]) UnmarshalJSON(data []byte) error {
	tree := jsonTree[T]{}
	err := json.Unmarshal(data, &tree)
	if err != nil {
		return err
	}
	t.fromJsonTree(tree)
	return nil
}

// UnmarshalJSON unmarshal data with parents
func (t *Tree[T]) fromJsonTree(tree jsonTree[T]) {
	t.Id = tree.Id
	t.Data = tree.Data
	list := make([]Tree[T], len(tree.Children))
	t.Children = make(TreeList[T], len(tree.Children))
	for i := range list {
		list[i].fromJsonTree(*tree.Children[i])
		list[i].Parent = t
		t.Children[i] = &list[i]
	}
	return
}

// Map tree items
func (t *Tree[T]) Map(c func(i *Tree[T])) {
	c(t)
	for i := range t.Children {
		t.Children[i].Map(c)
	}
}

// FirstChild get first child
func (t *Tree[T]) FirstChild() *Tree[T] {
	return t.Children.First()
}

// LastChild get last child
func (t *Tree[T]) LastChild() *Tree[T] {
	return t.Children.Last()
}

// GetById get by identifier
func (t *Tree[T]) GetById(id int) *Tree[T] {
	if t.Id == id {
		return t
	}
	for i := range t.Children {
		item := t.Children[i].GetById(id)
		if item != nil {
			return item
		}
	}
	return nil
}

// Path build path
func (t *Tree[T]) Path(delimiter string, c func(i *Tree[T]) string) string {
	if t.Parent != nil {
		return t.Parent.Path(delimiter, c) + delimiter + c(t)
	}
	return c(t)
}

// Lookup function
func (t *Tree[T]) Lookup(found func(i *Tree[T]) bool) TreeList[T] {
	var result = make(TreeList[T], 0, 4)
	if found(t) {
		result = append(result, t)
	}
	if len(t.Children) > 0 {
		result = append(result, t.Children.Lookup(found)...)
	}
	return result
}
