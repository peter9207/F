package average_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/peter9207/F/S/average"
)

var _ = Describe("Average", func() {

	data := []int64{1, 2, 3, 4, 5}

	Describe("Sum", func() {

		result := Sum(data)
		It("Should compute the sum", func() {
			Ω(result).Should(Equal(int64(15)))
		})

	})

})
