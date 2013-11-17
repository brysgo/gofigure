package decision_tree

import (
	bt "github.com/brysgo/gofigure/binary_tree"
	il "github.com/brysgo/gofigure/input_log"
)

type DecisionTree struct {
	bt.Tree
}
type Interface interface {
	bt.Interface
	Train(il.Interface, string)
	New() Interface
}

func (self *DecisionTree) New() Interface {
	t := DecisionTree{}
	return &t
}

func (self *DecisionTree) Train(inputLog il.Interface, targetKey string) {
	nextKey := inputLog.MaxEntropy(targetKey)
	if nextKey != targetKey {
		leftInputLog, rightInputLog := inputLog.SplitOnKey(nextKey)
		left, right := new(DecisionTree), new(DecisionTree)
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
		self.Set(inputContainsKey(inputLog.At(0), targetKey))
	}
}

func inputContainsKey(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
