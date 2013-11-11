package decision_tree

import (
  bt "github.com/brysgo/gofigure/binary_tree"
  il "github.com/brysgo/gofigure/input_log"
)

type DecisionTree struct {
  bt.Tree
}

func (self *DecisionTree) Train(inputLog il.InputLog, targetKey string) {
  nextKey := inputLog.MaxEntropy(targetKey)
  if nextKey != targetKey {
    leftInputLog, rightInputLog := inputLog.SplitOnKey(nextKey)
    left, right := new(DecisionTree), new(DecisionTree)
    if len(leftInputLog) > 0 {
      left.Train(leftInputLog, targetKey)
    }
    if len(rightInputLog) > 0 {
      right.Train(rightInputLog, targetKey)
    }
    self.Left, self.Right = &left.Tree, &right.Tree
    self.Value = nextKey
  } else {
    self.Value = inputLog[0].HasKey(targetKey)
  }
}

