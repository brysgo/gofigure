package gofigure_test

import (
  dt "../lib/decision_tree"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func init() {
  var _ = Describe("Decision Tree", func() {

    It("builds a decision tree from a list of input features and an output feature", func() {
      xorInputSet := dt.InputLog{
        {"0":true, "1":false, "2":true},
        {"0":false,"1":false,"2":false},
        {"0":true, "1":true, "2":false},
        {"0":false, "1":true, "2":true},
      }

      myTree := new(dt.DecisionTree)
      myTree.Train(xorInputSet, "2")

      Expect(myTree.Left).ToNot(BeNil())
      Expect(myTree.Left.Right).ToNot(BeNil())
      Expect(myTree.Left.Right.Value).To(Equal(true))
    })

    It("it uses ID3 to maximize entropy", func(){
      id3InputSet := dt.InputLog{
        {"0": true,"1":false,"2":false,"3": true},
        {"0": true,"1": true,"2":false,"3": true},
        {"0":false,"1":false,"2":false,"3": true},
        {"0":false,"1": true,"2":false,"3": true},
        {"0": true,"1":false,"2": true,"3":false},
        {"0": true,"1": true,"2": true,"3":false},
        {"0":false,"1":false,"2": true,"3":false},
        {"0":false,"1": true,"2": true,"3":false},
      }

      myTree := new(dt.DecisionTree)
      myTree.Train(id3InputSet, "3")

      Expect(myTree.Value).To(Equal("2"))
    })

    XIt("doesn't continue building a branch if a decision can be made", func(){})

  })
}
