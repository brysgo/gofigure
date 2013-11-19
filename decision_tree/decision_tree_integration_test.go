package decision_tree_test

import (
  dt "github.com/brysgo/gofigure/decision_tree"
  il "github.com/brysgo/gofigure/input_log"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("DecisionTree", func() {

  Describe("Training a Decision Tree", func() {
    var id3InputSet il.InputLog
    var myTree dt.Interface

    BeforeEach(func() {
      id3InputSet = il.InputLog{
        {"1", "3"},
        {"0", "1", "3"},
        {"1", "3"},
        {"0", "1", "3"},
        {"1", "2"},
        {"0", "1", "2"},
        {},
        {"0"},
      }

      myTree = dt.New()
      myTree.Train(il.Interface(id3InputSet), "3")
    })

    It("it uses ID3 to maximize entropy", func() {
      Expect(myTree.Value()).To(Equal("2"))
      Expect(myTree.Left().Value()).To(Equal("1"))
    })

    It("doesn't continue building a branch if a decision can be made", func() {
      Expect(myTree.Left().Right().Right()).To(BeNil())
    })

    It("builds a tree with the right decisions", func() {
      Expect(myTree.Left().Right().Value()).To(Equal(true))
      Expect(myTree.Left().Left().Value()).To(Equal(false))
      Expect(myTree.Right().Value()).To(Equal(false))
    })
  })

})
