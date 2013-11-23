package decision_tree

import (
  bt "github.com/brysgo/gofigure/binary_tree"
  il "github.com/brysgo/gofigure/input_log"
)

type DecisionTree struct {
  bt.Interface
}
type Interface interface {
  bt.Interface
  Train(il.Interface, string)
  Decide([]string) bool
}

var New = func() Interface {
  t := DecisionTree{
    Interface: bt.New(),
  }
  return &t
}

func (self *DecisionTree) Train(inputLog il.Interface, targetKey string) {
  nextKey := inputLog.MaxEntropy(targetKey)
  if nextKey != targetKey {
    leftInputLog, rightInputLog := inputLog.SplitOnKey(nextKey)
    left, right := New(), New()
    if leftInputLog.Any() {
      left.Train(leftInputLog, targetKey)
    }
    if rightInputLog.Any() {
      right.Train(rightInputLog, targetKey)
    }
    self.InsertLeft(left)
    self.InsertRight(right)
    self.Set(nextKey)
  } else {
    present := inputLog.KeyCount(targetKey)
    absent := inputLog.Len() - present
    self.Set(present > absent)
  }
}

func (self *DecisionTree) Decide(input []string) bool {
  nodeLabel := self.Value()

  switch nodeLabel.(type) {
  case bool:
    return nodeLabel.(bool)
  default:
    nodeLabelInInput := false
    for _, key := range input {
      if nodeLabel == key {
        nodeLabelInInput = true
        break
      }
    }
    var nextNode Interface
    if nodeLabelInInput {
      nextNode = self.Right().(Interface)
    } else {
      nextNode = self.Left().(Interface)
    }
    return nextNode.Decide(input)
  }

}
