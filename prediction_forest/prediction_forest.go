package prediction_forest

import (
	dt "github.com/brysgo/gofigure/decision_tree"
	il "github.com/brysgo/gofigure/input_log"
)

type PredictionForest struct {
	Trees        map[string]dt.Interface
	treeTemplate dt.Interface
}

func New(treeTemplate dt.Interface) *PredictionForest {
	forest := PredictionForest{
		Trees:        make(map[string]dt.Interface),
		treeTemplate: treeTemplate,
	}
	return &forest
}

func (self *PredictionForest) Train(inputLog il.InputLog) {
	for _, key := range inputLog.Keys() {
		var newTree dt.Interface
		newTree = self.treeTemplate.New()
		newTree.Train(inputLog, key)
		self.Trees[key] = newTree
	}
}
