package prediction_forest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrediction_forest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prediction_forest Suite")
}
