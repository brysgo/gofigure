package gofigure_test

import (
  bt "../lib/binary_tree"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func init() {
  var _ = Describe("Binary Tree", func() {

    It("creates a new empty tree", func() {
      var t *bt.Tree = bt.New()

      Expect(t.Value).To(BeNil())
    })

    It("allows insertions with an arbitrary value", func() {
      var t *bt.Tree = bt.New()

      t.Value = 5

      type ExampleType struct {
        Arbitrary int
        Fields string
      }

      ex := ExampleType{ }
      node := t.InsertLeft(ex)

      Expect(node).To(Equal(t.Left))
      Expect(node.Value).To(Equal(ex))

      right_node := node.InsertRight("hello")

      Expect(right_node).To(Equal(node.Right))
      Expect(right_node.Value).To(Equal("hello"))
    })

  })
}
