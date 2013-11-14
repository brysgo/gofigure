package tag_soup_test

import (
	ts "github.com/brysgo/gofigure/tag_soup"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TagSoup", func() {

	Describe("New", func() {
		It("returns a new tag", func() {
			Expect(ts.New()).ToNot(BeNil())
		})
	})

	Describe("AddTagGroup", func() {
		var exampleTagGroup ts.TagGroup
		var tagSoup, soupPointer *ts.TagSoup

		BeforeEach(func() {
			exampleTagGroup = ts.TagGroup{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
			tagSoup = ts.New()
			soupPointer = tagSoup.AddTagGroup(exampleTagGroup)
		})

		It("returns a pointer the receiver", func() {
			Expect(soupPointer).To(Equal(tagSoup))
		})

		It("adds the tag group to the input log", func() {
			Expect(len(tagSoup.InputLog)).To(Equal(1))
		})

		PIt("trains a bunch of decision trees to predict tags", func() {
		})
	})

})
