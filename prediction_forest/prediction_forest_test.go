package prediction_forest_test

import (
  "code.google.com/p/gomock/gomock"
  dt "github.com/brysgo/gofigure/decision_tree"
  il "github.com/brysgo/gofigure/input_log"
  pf "github.com/brysgo/gofigure/prediction_forest"
  . "github.com/onsi/ginkgo"
  "github.com/onsi/ginkgo/thirdparty/gomocktestreporter"
  . "github.com/onsi/gomega"
  mdt "gofigure_mocks/mock_decision_tree"
)

var _ = Describe("PredictionForest", func() {

  var (
    mockCtrl        *gomock.Controller
    myForest        *pf.PredictionForest
    decisionTreeNew func() dt.Interface
  )

  BeforeEach(func() {
    decisionTreeNew = dt.New
    mockCtrl = gomock.NewController(gomocktestreporter.New())
  })

  AfterEach(func() {
    dt.New = decisionTreeNew
    mockCtrl.Finish()
  })

  Describe("Train", func() {
    var (
      predictionForestInputSet il.InputLog
      mockDecisionTree         *mdt.MockInterface
    )

    BeforeEach(func() {
      mockDecisionTree = mdt.NewMockInterface(mockCtrl)
      dt.New = func() dt.Interface {
        return mockDecisionTree
      }
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
        mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
        mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
        mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
        mockDecisionTree.EXPECT().Train(gomock.Any(), gomock.Any()),
      )
      myForest = pf.New(mockDecisionTree)
      myForest.Train(predictionForestInputSet)
    })

  })

  Describe("Predict", func() {
    var (
      mockHelloTree, mockWorldTree *mdt.MockInterface
    )

    BeforeEach(func() {
      mockHelloTree = mdt.NewMockInterface(mockCtrl)
      mockWorldTree = mdt.NewMockInterface(mockCtrl)
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
