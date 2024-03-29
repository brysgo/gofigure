package prediction_forest

import (
  dt "github.com/brysgo/gofigure/decision_tree"
  il "github.com/brysgo/gofigure/input_log"
)

type PredictionForest struct {
  Trees        map[string]dt.Interface
}
type Interface interface {
	Train(il.Interface)
	Predict([]string) []string
}

func New() Interface {
  forest := PredictionForest{
    Trees:        make(map[string]dt.Interface),
  }
  return &forest
}

func (self *PredictionForest) Train(inputLog il.Interface) {
  for _, key := range inputLog.Keys() {
    var newTree dt.Interface
    newTree = dt.New()
    newTree.Train(inputLog, key)
    self.Trees[key] = newTree
  }
}

func (self *PredictionForest) Predict(input []string) (result []string) {
  for key, tree := range self.Trees {
    if tree.Decide(input) {
      result = append(result, key)
    }
  }
  return
}
