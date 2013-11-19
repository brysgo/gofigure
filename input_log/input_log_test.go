package input_log_test

import (
  il "github.com/brysgo/gofigure/input_log"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("InputLog", func() {
  var (
    data il.InputLog
  )

  BeforeEach(func() {
    data = il.InputLog{
      {"0", "1", "2", "3"},
      {"0", "1", "2", "3"},
      {"0", "1", "2", "3"},
      {"0", "1", "2", "3"},
      {"0", "1", "2"},
      {"0", "1", "2"},
      {"0", "1"},
      {"0", "1"},
      {"0", "1"},
      {"0"},
      {"0"},
      {"0"},
      {"0"},
    }
  })

  Describe("MaxEntropy", func() {

    It("returns the key that best decides on the target", func() {
      Expect(data.MaxEntropy("3")).To(Equal("2"))
    })

  })

  Describe("SplitOnKey", func() {

    It("returns two intput logs, one that had the key, one that did not", func() {
      left, right := data.SplitOnKey("2")
      Expect(left.Len()).To(Equal(7))
      Expect(right.Len()).To(Equal(6))
    })

  })

  Describe("AddInput", func() {

    It("appends an input to the input log", func() {
      newInput := []string{"4", "5"}
      data = data.AddInput(newInput).(il.InputLog)
      Expect(data[len(data)-1]).To(Equal(newInput))
    })

  })

  Describe("Keys", func() {

    It("returns a complete list of the input log's strings", func() {
      Expect(data.Keys()).To(Equal([]string{"0", "1", "2", "3"}))
    })

  })

  Describe("KeyCount", func() {

    It("returns the number of key occurrances in the input log", func() {
      Expect(data.KeyCount("1")).To(Equal(9))
    })
  })

})
