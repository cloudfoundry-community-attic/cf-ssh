package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCFSSH(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cfssh suite")
}
