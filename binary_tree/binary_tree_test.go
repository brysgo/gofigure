package binary_tree_test

import (
  bt "github.com/brysgo/gofigure/binary_tree"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("BinaryTree", func() {

  It("creates a new empty tree", func() {
    var t *bt.Tree = bt.New()

    Expect(t.Value()).To(BeNil())
  })

  It("allows insertions with an arbitrary value", func() {
    var t *bt.Tree = bt.New()

    t.Set(5)

    type ExampleType struct {
      Arbitrary int
      Fields    string
    }

    ex := ExampleType{}
    child := bt.New()
    child.Set(ex)
    node := t.InsertLeft(child)

    Expect(node).To(Equal(t.Left()))
    Expect(node.Value()).To(Equal(ex))

    child = bt.New()
    child.Set("hello")
    rightNode := node.InsertRight(child)

    Expect(rightNode).To(Equal(node.Right()))
    Expect(rightNode.Value()).To(Equal("hello"))
  })
})
