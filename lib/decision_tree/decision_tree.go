package gofigure

import (
  bt "../binary_tree"
)

type DecisionTree struct {
  bt.Tree
}
type Input map[string]bool
type InputLog []Input

func (self *DecisionTree) Train(inputLog InputLog, targetKey string) {
  nextKey := inputLog.maxEntropy(targetKey)
  if nextKey != targetKey {
    leftInputLog, rightInputLog := inputLog.splitOnKey(nextKey)
    left, right := new(DecisionTree), new(DecisionTree)
    left.Train(leftInputLog, targetKey)
    right.Train(rightInputLog, targetKey)
    self.Left, self.Right = &left.Tree, &right.Tree
    self.Value = nextKey
  } else {
    self.Value = inputLog[0][targetKey]
  }
}

func (self InputLog) maxEntropy(targetKey string) string {
  entropy := make(map[string]int, len(self[0])-1)
  for _, input := range self {
    for key, value := range input {
      if key != targetKey {
        if value == input[targetKey] {
          entropy[key] += 1
        } else {
          entropy[key] += 0
        }
      }
    }
  }
  maxEntropy := 0
  resultKey := targetKey
  for key, value := range entropy {
    ent := value - len(self)/2
    if ent < 0 { ent *= -1 }
    if ent >= maxEntropy {
      maxEntropy = ent
      resultKey = key
    }
  }
  return resultKey
}

func (self InputLog) splitOnKey(targetKey string) (left InputLog, right InputLog) {
  for _, v := range self {
    remainder := make(Input, len(v)-1)
    for k, b := range v {
      if k != targetKey {
        remainder[k]=b
      }
    }
    if v[targetKey] {
      right = append(right, remainder)
    } else {
      left = append(left, remainder)
    }
  }
  return
}

