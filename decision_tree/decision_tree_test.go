package decision_tree_test

import (
  "code.google.com/p/gomock/gomock"
  bt "github.com/brysgo/gofigure/binary_tree"
  dt "github.com/brysgo/gofigure/decision_tree"
  . "github.com/onsi/ginkgo"
  "github.com/onsi/ginkgo/thirdparty/gomocktestreporter"
  mbt "gofigure_mocks/mock_binary_tree"
  mdt "gofigure_mocks/mock_decision_tree"
  mil "gofigure_mocks/mock_input_log"
)

var (
  mockCtrl        *gomock.Controller
  myTree          dt.Interface
  mockBinaryTree  *mbt.MockInterface
  binaryTreeNew   func() bt.Interface = bt.New
  decisionTreeNew func() dt.Interface = dt.New
)

var _ = Describe("DecisionTree", func() {
  BeforeEach(func() {
    mockCtrl = gomock.NewController(gomocktestreporter.New())
    mockBinaryTree = mbt.NewMockInterface(mockCtrl)
    bt.New = func() bt.Interface {
      return mockBinaryTree
    }
    myTree = dt.New()
  })

  AfterEach(func() {
    mockCtrl.Finish()

    bt.New = binaryTreeNew
    dt.New = decisionTreeNew
  })

  Describe("#Train", func() {
    var (
      inputLog, leftInputLog, rightInputLog *mil.MockInterface
      targetKey, nextKey                    string
      maxEntropyCall                        *gomock.Call
      keyCountCall                          *gomock.Call
      leftDecisionTree, rightDecisionTree   *mdt.MockInterface
      keySetCall, leftAnyCall, rightAnyCall *gomock.Call
    )

    BeforeEach(func() {
      inputLog = mil.NewMockInterface(mockCtrl)
      targetKey = "the target key"

      leftInputLog = mil.NewMockInterface(mockCtrl)
      rightInputLog = mil.NewMockInterface(mockCtrl)

      maxEntropyCall = inputLog.EXPECT().MaxEntropy(targetKey)
    })

    Describe("when there are no features left to decide on", func() {

      BeforeEach(func() {
        maxEntropyCall.Return(targetKey)
        inputLog.EXPECT().Len().Return(7)
        keyCountCall = inputLog.EXPECT().KeyCount(targetKey)
      })

      Context("when there are more occurances of the target key than absences", func() {
        It("decides with a value of true", func() {
          keyCountCall.Return(5)
          mockBinaryTree.EXPECT().Set(true)
          myTree.Train(inputLog, targetKey)
        })
      })
      Context("when there are more absences of the target key than occurances", func() {
        It("decides with a value of false", func() {
          keyCountCall.Return(3)
          mockBinaryTree.EXPECT().Set(false)
          myTree.Train(inputLog, targetKey)
        })
      })
    })

    Describe("when there are still features to decide on", func() {

      BeforeEach(func() {
        nextKey = "the next key to decide on"
        maxEntropyCall = maxEntropyCall.Return(nextKey)

        leftDecisionTree = mdt.NewMockInterface(mockCtrl)
        rightDecisionTree = mdt.NewMockInterface(mockCtrl)

        decisionTreeMocks := []*mdt.MockInterface{rightDecisionTree, leftDecisionTree}
        // Inject left and right tree in alternating fashion
        leftRightParity := 0
        dt.New = func() dt.Interface {
          leftRightParity++
          return decisionTreeMocks[leftRightParity%2]
        }

        inputLog.EXPECT().SplitOnKey(nextKey).Return(leftInputLog, rightInputLog)
        mockBinaryTree.EXPECT().InsertLeft(leftDecisionTree)
        mockBinaryTree.EXPECT().InsertRight(rightDecisionTree)

        keySetCall = mockBinaryTree.EXPECT().Set(nextKey).AnyTimes()
        rightAnyCall = rightInputLog.EXPECT().Any().Return(false)
        leftAnyCall = leftInputLog.EXPECT().Any().Return(false)
      })

      Context("and the deciding feature has positive entries", func() {

        BeforeEach(func() {
          rightAnyCall.Return(true)
        })

        It("recurses on the right input set", func() {
          rightDecisionTree.EXPECT().Train(rightInputLog, targetKey)
          myTree.Train(inputLog, targetKey)
        })

      })

      Context("and the deciding feature has negative entries", func() {

        BeforeEach(func() {
          leftAnyCall.Return(true)
        })

        It("recurses on the left input set", func() {
          leftDecisionTree.EXPECT().Train(leftInputLog, targetKey)
          myTree.Train(inputLog, targetKey)
        })

      })

      It("assigns the key of the deciding feature to the current node", func() {
        keySetCall.Times(1)
        myTree.Train(inputLog, targetKey)
      })
    })

  })

  Describe("#Decide", func() {
  })

})
