//go:build provider_new
// +build provider_new

package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestOnlyProviderDemo(realT *testing.T) {
	r := runner.NewRunner(realT, realT.Name())
	r.BeforeEach(func(t provider.T) {
		t.Epic("Only Provider Demo")
		t.Feature("runner.RunTest")
		t.Description("`allure-go allows you to use allure without suites. Even if it broken or failed`")
	})
	r.NewTest("Broken test", func(t provider.T) {
		t.Title("Some broken test")
		panic("whoops")
	})

	r.NewTest("Failed test", func(t provider.T) {
		t.Title("Some failed test")
		t.Require().NotNil(nil)
	})

	r.NewTest("Passed test", func(t provider.T) {
		t.Title("Some passed test")
	})

	r.RunTests()
}
