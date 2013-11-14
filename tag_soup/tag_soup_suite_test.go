package tag_soup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSmarty_pants(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tag Soup Suite")
}
