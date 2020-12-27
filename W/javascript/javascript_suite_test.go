package javascript_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJavascript(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Javascript Suite")
}
