package prediction_forest_test

import (
	gomock "code.google.com/p/gomock/gomock"
	il "github.com/brysgo/gofigure/input_log"
	dt "github.com/brysgo/gofigure/mocks/mock_decision_tree"
	pf "github.com/brysgo/gofigure/prediction_forest"
	gomock_ginkgo "github.com/brysgo/gomock_ginkgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PredictionForest", func() {

	var (
		mockCtrl         *gomock.Controller
		mockDecisionTree *dt.MockInterface
		myForest         *pf.PredictionForest
	)

	BeforeEach(func() {
		mockCtrl = gomock_ginkgo.NewController()
		mockDecisionTree = dt.NewMockInterface(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Train", func() {
		var (
			predictionForestInputSet il.InputLog
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

		})

		It("Builds a map of decision trees targeted at each input", func() {
			gomock.InOrder(
				mockDecisionTree.EXPECT().New().Return(mockDecisionTree),
				mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
				mockDecisionTree.EXPECT().New().Return(mockDecisionTree),
				mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
				mockDecisionTree.EXPECT().New().Return(mockDecisionTree),
				mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
				mockDecisionTree.EXPECT().New().Return(mockDecisionTree),
				mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
			)
			myForest = pf.New(mockDecisionTree)
			myForest.Train(predictionForestInputSet)
			Expect(true).ToNot(Equal(false))
			// Expect(myForest.Trees["0"]).ToNot(Equal(myForest.Trees["1"]))
		})

	})

	Describe("ExpectedSet", func() {})

})
