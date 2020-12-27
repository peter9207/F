package javascript_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/F/W/javascript"
)

var _ = Describe("Javascript", func() {

	Describe("Depth", func() {

		When("root is a simple number", func() {

			root := 1
			It("returns 1", func() {
				result := javascript.Depth(root)
				Expect(result).Should(Equal(1))
			})
		})

	})

})
