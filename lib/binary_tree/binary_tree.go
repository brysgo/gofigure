package bt

type AnyType interface{}
type Comparer func(one, two AnyType) (bool)

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

func (parent *Tree) Insert(value AnyType, lessThan Comparer) (child *Tree) {
  
  if lessThan(value, parent.Value) {
    if parent.Left == nil {
      child = parent.InsertLeft(value)
    } else {
      child = parent.Left.Insert(value, lessThan)
    }
  } else {
    if parent.Right == nil {
      child = parent.InsertRight(value)
    } else {
      child = parent.Right.Insert(value, lessThan)
    }
  }

  return
}
