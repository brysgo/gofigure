package decision_tree_test

import (
  dt "../decision_tree"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("DecisionTree", func() {

  Describe("Training a Decision Tree", func() {
    var id3InputSet dt.InputLog
    var myTree *dt.DecisionTree

    BeforeEach( func() {
      id3InputSet = dt.InputLog{
        {"0":false,"1": true,"2":false,"3": true},
        {"0": true,"1": true,"2":false,"3": true},
        {"0":false,"1": true,"2":false,"3": true},
        {"0": true,"1": true,"2":false,"3": true},
        {"0":false,"1": true,"2": true,"3":false},
        {"0": true,"1": true,"2": true,"3":false},
        {"0":false,"1":false,"2":false,"3":false},
        {"0": true,"1":false,"2":false,"3":false},
      }

      myTree = new(dt.DecisionTree)
      myTree.Train(id3InputSet, "3")
    })

    It("it uses ID3 to maximize entropy", func(){
      Expect(myTree.Value).To(Equal("2"))
      Expect(myTree.Left.Value).To(Equal("1"))
    })

    It("doesn't continue building a branch if a decision can be made", func(){
      Expect(myTree.Left.Right.Right).To(BeNil())
    })

    It("builds a tree with the right decisions", func() {
      Expect(myTree.Left.Right.Value).To(Equal(true))
      Expect(myTree.Left.Left.Value).To(Equal(false))
      Expect(myTree.Right.Value).To(Equal(false))
    })
  })

})
