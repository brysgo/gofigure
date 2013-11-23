package prediction_forest_test

import (
  il "github.com/brysgo/gofigure/input_log"
  pf "github.com/brysgo/gofigure/prediction_forest"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("PredictionForest", func() {

  var inputSet il.InputLog
  var myForest pf.Interface

  BeforeEach(func() {
    inputSet = il.InputLog{
      {},
      {"0", "1", "3"},
      {"1", "3"},
      {"0", "3"},
      {"0", "1", "2", "3"},
      {"0", "1", "3"},
      {"0", "1"},
      {"0", "2", "4"},
      {"2", "4"},
      {"4", "1"},
      {"0", "4", "2"},
    }

    myForest = pf.New()
    myForest.Train(il.Interface(inputSet))
  })

  Describe("Making a prediction", func() {
    It("adds a key that should be there", func() {
      Expect(myForest.Predict([]string{"0", "1"})).To(Equal([]string{"0", "1", "3"}))
    })

    It("removes a key that doesn't belong", func() {
      Expect(myForest.Predict([]string{"0", "1", "2", "3"})).To(Equal([]string{"0", "1", "3"}))
    })

    It("finds keys associated with a key", func() {
      Expect(myForest.Predict([]string{"4"})).To(Equal([]string{"0", "1", "2"}))
    })
  })

})
