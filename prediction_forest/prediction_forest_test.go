package prediction_forest_test

import (
	il "github.com/brysgo/gofigure/input_log"
	pf "github.com/brysgo/gofigure/prediction_forest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PredictionForest", func() {

	Describe("Train", func() {
		var (
			predictionForestInputSet il.InputLog
			myForest                 *pf.PredictionForest
		)

		BeforeEach(func() {
			predictionForestInputSet = il.InputLog{
				{"1", "3"},
				{"0", "1", "3"},
				{"1", "3"},
				{"0", "1", "3"},
				{"1", "2"},
				{"0", "1", "2"},
				{},
				{"0"},
			}

			myForest = pf.New()
			myForest.Train(predictionForestInputSet)
		})

		It("Builds a map of decision trees targeted at each input", func() {
			Expect(len(*myForest)).To(Equal(4))
			Expect((*myForest)["0"].Value).ToNot(BeNil())
			Expect((*myForest)["1"].Value).ToNot(BeNil())
			Expect((*myForest)["2"].Value).ToNot(BeNil())
			Expect((*myForest)["3"].Value).ToNot(BeNil())
		})

	})

	Describe("ExpectedSet", func() {})

})
