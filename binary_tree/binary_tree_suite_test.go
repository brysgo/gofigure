package binary_tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBinary_tree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Binary_tree Suite")
}
