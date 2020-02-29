package json_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJson(t *testing.T) {
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "JSON Suite")
}
