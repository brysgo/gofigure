package binary_tree_test

import (
	bt "../binary_tree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinaryTree", func() {

	It("creates a new empty tree", func() {
		var t *bt.Tree = bt.New()

		Expect(t.Value).To(BeNil())
	})

	It("allows insertions with an arbitrary value", func() {
		var t *bt.Tree = bt.New()

		t.Value = 5

		type ExampleType struct {
			Arbitrary int
			Fields    string
		}

		ex := ExampleType{}
		node := t.InsertLeft(ex)

		Expect(node).To(Equal(t.Left))
		Expect(node.Value).To(Equal(ex))

		rightNode := node.InsertRight("hello")

		Expect(rightNode).To(Equal(node.Right))
		Expect(rightNode.Value).To(Equal("hello"))
	})

	It("inserts values in the right order", func() {
		var t *bt.Tree = bt.New()

		t.Value = 5

		compare := bt.Comparer(func(one, two bt.AnyType) bool {
			return one.(int) < two.(int)
		})

		node := t.Insert(3, compare)

		Expect(node).To(Equal(t.Left))
		Expect(node.Value).To(Equal(3))

		rightNode := node.Insert(4, compare)

		Expect(rightNode).To(Equal(node.Right))
		Expect(rightNode.Value).To(Equal(4))
	})

})
