package shapes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/F/S/shapes"
)

var _ = Describe("LinkedList", func() {

	Describe("isEmpty", func() {
		list := shapes.NewLinkedList()
		It("should be empty", func() {
			actual := list.IsEmpty()
			Ω(actual).Should(Equal(true))
		})

		It("should not be empty", func() {
			list.AddFirst(25)
			actual := list.IsEmpty()
			Ω(actual).Should(Equal(false))
		})

	})

	Describe("adding and removing", func() {
		list := shapes.NewLinkedList()
		list.AddFirst(1)
		list.AddFirst(2)

		Describe("Rmoeve First", func() {
			It("should return 2", func() {
				actual := list.RemoveFirst()
				Ω(actual).Should(Equal(float64(2)))
			})
		})

	})

})
