package input_log_test

import (
	il "github.com/brysgo/gofigure/input_log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InputLog", func() {
	var data il.InputLog

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
			Expect(len(left)).To(Equal(7))
			Expect(len(right)).To(Equal(6))
		})

	})

})
