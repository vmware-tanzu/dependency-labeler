package main_test

import (
	"github.com/onsi/gomega/gexec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	pathToBin string
)

func TestDeplab(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		var (
			err error
		)

		pathToBin, err = gexec.Build("github.com/pivotal/deplab")
		Expect(err).ToNot(HaveOccurred())
	})

	RunSpecs(t, "Deplab Suite")
}
