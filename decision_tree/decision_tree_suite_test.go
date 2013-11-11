package decision_tree_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"
)

func TestDecision_tree(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Decision_tree Suite")
}
