package prediction_forest

import (
	dt "github.com/brysgo/gofigure/decision_tree"
	il "github.com/brysgo/gofigure/input_log"
)

type PredictionForest map[string]*dt.DecisionTree

func New() *PredictionForest {
	forest := make(PredictionForest)
	return &forest
}

func (self *PredictionForest) Train(inputLog il.InputLog) {
	for _, key := range inputLog.Keys() {
		decisionTree := new(dt.DecisionTree)
		decisionTree.Train(inputLog, key)
		(*self)[key] = decisionTree
	}
}
