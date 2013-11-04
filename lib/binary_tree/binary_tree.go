package gofigure

type AnyType interface{}

type Tree struct {
  Left  *Tree
  Value AnyType
  Right *Tree
}

func New() *Tree {
  t := Tree{}
  return &t
}

func (parent *Tree) InsertLeft(value AnyType) (child *Tree) {
  child = New()
  child.Value = value
  parent.Left = child
  return
}

func (parent *Tree) InsertRight(value AnyType) (child *Tree) {
  child = New()
  child.Value = value
  parent.Right = child
  return
}
