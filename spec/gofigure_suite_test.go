package gofigure_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"
)

func TestGofigure(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Gofigure")
}
