package decision_tree

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
    if len(leftInputLog) > 0 {
      left.Train(leftInputLog, targetKey)
    }
    if len(rightInputLog) > 0 {
      right.Train(rightInputLog, targetKey)
    }
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
      if key == targetKey { 
        if value {
          entropy[targetKey] += 1
        }
        continue
      }
      if value == input[targetKey] {
        entropy[key] += 1
      } else {
        entropy[key] += 0
      }
    }
  }
  if entropy[targetKey] % len(self) == 0 { return targetKey }
  maxEntropy := 0.0
  resultKey := targetKey
  for key, value := range entropy {
    ent := float64(value - len(self)/2)
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

