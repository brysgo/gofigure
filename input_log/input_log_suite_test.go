package input_log_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"
)

func TestInput_log(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Input_log Suite")
}
