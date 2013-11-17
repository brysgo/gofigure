package binary_tree

type AnyType interface{}

type Tree struct {
  left  Interface
  value AnyType
  right Interface
}
type Interface interface {
  InsertLeft(Interface) Interface
  InsertRight(Interface) Interface
  Left() Interface
  Right() Interface
  Value() AnyType
  Set(AnyType)
}

func New() Interface {
  t := Tree{}
  return &t
}

func (parent *Tree) InsertLeft(node Interface) Interface {
  parent.left = node
  return node
}

func (parent *Tree) InsertRight(node Interface) Interface {
  parent.right = node
  return node
}

func (self *Tree) Left() Interface {
  return self.left
}

func (self *Tree) Right() Interface {
  return self.right
}

func (self *Tree) Value() AnyType {
  return self.value
}

func (self *Tree) Set(value AnyType) {
  self.value = value
}
