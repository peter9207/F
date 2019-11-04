package average_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/peter9207/F/S/average"
)

var _ = Describe("Average", func() {

	data := []int64{1, 2, 1, 2, 1, 2}

	Describe("Sum", func() {

		result := Sum(data)
		It("Should compute the sum", func() {
			Ω(result).Should(Equal(int64(9)))
		})

	})

	Describe("Rolling average", func() {

		averages := Rolling(data, 2)

		It("should have a length equal to data", func() {
			Ω(len(averages)).Should(Equal(6))
		})

		It("should have correct values", func() {
			for _, v := range averages {
				Ω(v).Should(Equal(1.5))
			}
		})

	})

})
