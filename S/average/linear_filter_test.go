package average_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/F/S/average"
)

var _ = Describe("LinearFilter", func() {

	data := []float64{1, 2, 1, 2, 1, 2}

	result := average.Box(data, 2)

	It("should have a length equal to data", func() {
		Ω(len(result)).Should(Equal(len(data)))
	})
})
