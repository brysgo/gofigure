package prediction_forest_test

import (
  "code.google.com/p/gomock/gomock"
  il "github.com/brysgo/gofigure/input_log"
  pf "github.com/brysgo/gofigure/prediction_forest"
  . "github.com/onsi/ginkgo"
  "github.com/onsi/ginkgo/thirdparty/gomocktestreporter"
  . "github.com/onsi/gomega"
  dt "gofigure_mocks/mock_decision_tree"
)

var _ = Describe("PredictionForest", func() {

  var (
    mockCtrl *gomock.Controller
    myForest *pf.PredictionForest
  )

  BeforeEach(func() {
    mockCtrl = gomock.NewController(gomocktestreporter.New())
  })

  AfterEach(func() {
    mockCtrl.Finish()
  })

  Describe("Train", func() {
    var (
      predictionForestInputSet il.InputLog
      mockDecisionTree         *dt.MockInterface
    )

    BeforeEach(func() {
      mockDecisionTree = dt.NewMockInterface(mockCtrl)
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
    })

  })

  Describe("Predict", func() {
    var (
      mockHelloTree, mockWorldTree *dt.MockInterface
    )

    BeforeEach(func() {
      mockHelloTree = dt.NewMockInterface(mockCtrl)
      mockWorldTree = dt.NewMockInterface(mockCtrl)
      myForest = pf.New(mockHelloTree)
      myForest.Trees["hello"] = mockHelloTree
      myForest.Trees["world"] = mockWorldTree
    })

    It("Returns the result of running all the decision trees", func() {
      input := []string{"world", "dog", "cat"}
      output := []string{"hello"}
      mockHelloTree.EXPECT().Decide(input).Return(true)
      mockWorldTree.EXPECT().Decide(input).Return(false)
      Expect(myForest.Predict(input)).To(Equal(output))
    })

  })

})
