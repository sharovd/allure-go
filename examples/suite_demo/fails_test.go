//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type FailsDemoSuite struct {
	suite.Suite
}

func (s *FailsDemoSuite) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Failures")
}

func (s *FailsDemoSuite) TestAssertionFailMessage(t provider.T) {
	t.Title("This test failed with message by assert")
	t.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2
		Error message: Assertion Failed`)
	t.Tags("fail", "assertions")

	t.Require().Equal(1, 2, "Assertion Failed")
}

func (s *FailsDemoSuite) TestXSkip(t provider.T) {
	t.Title("This test skipped")
	t.Description(`
		This Test will be skipped`)
	t.Tags("fail", "xskip")

	t.XSkip()
	t.Require().Equal(1, 2, "Never reach this")
}

func (s *FailsDemoSuite) TestAssertionFailNoMessage(t provider.T) {
	t.Title("This test failed without message by assert")
	t.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2`)

	t.Tags("fail", "assertions")

	t.Require().Equal(1, 2)
}

func (s *FailsDemoSuite) TestAssertionFailInnerSteps(t provider.T) {
	t.Title("This test failed with message by assert in an inner step")
	t.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2
		Error message: Failed inside step`)

	t.Tags("fail", "assertions", "nesting")

	t.WithNewStep("Failed parent step", func(ctx provider.StepCtx) {
		ctx.WithNewStep("Failed child step", func(ctx provider.StepCtx) {
			ctx.Require().Equal(1, 2, "Failed inside step")
		})
	})
}

func (s *FailsDemoSuite) TestPanic(t provider.T) {
	t.Title("This test broken by panic")
	t.Description(`
		This Test will be broken by panic.
		Error text:
		test panicked: runtime error: index out of range [0] with length 0...`)

	t.Tags("fail", "broken", "panic")

	var test []string
	test2 := test[0]
	t.Require().Equal(test2, test2, "Never reach this")
}

func (s *FailsDemoSuite) TestPanicInnerSteps(t provider.T) {
	t.Title("This test broken by panic in an inner step")
	t.Description(`
		This Test will be broken by panic.
		All steps that includes panic will be broken.
		Error text:
		test panicked: runtime error: index out of range [0] with length 0...`)

	t.Tags("fail", "broken", "panic", "nesting")

	t.WithNewStep("Check 1", func(ctx provider.StepCtx) {
		ctx.WithNewStep("Check 1.1", func(ctx provider.StepCtx) {

		})
		ctx.WithNewStep("Check 1.2", func(ctx provider.StepCtx) {
			ctx.WithNewStep("Check 1.2.1", func(ctx provider.StepCtx) {
				var test []string
				test2 := test[0]
				ctx.Require().Equal(test2, test2, "Never reach this")
			})
		})
	})
}

func (s *FailsDemoSuite) TestBrokenStatusNoMessage(t provider.T) {
	t.Title("This test broken without message")
	t.Description(`
		This Test will be broken.
		No any message expected there`)

	t.Tags("fail", "broken")

	t.NewStep("This step will be reached before failing")
	t.Broken()
	t.NewStep("This step will be reached after failing")
}

func (s *FailsDemoSuite) TestBrokenNowStatusNoMessage(t provider.T) {
	t.Title("This test broken immediately without message")
	t.Description(`
		This Test will be broken.
		No any message expected there`)

	t.Tags("fail", "broken")

	t.NewStep("This step will be reached before failing")
	t.BrokenNow()
	t.NewStep("This step will never be reached after failing")
}

func (s *FailsDemoSuite) TestBrokenNowStatusMessage(t provider.T) {
	t.Title("This test broken immediately with message")
	t.Description(`
		This Test will be broken.
		Error message:
		Test fails as FailNow()`)

	t.Tags("fail", "broken")

	t.NewStep("This step will be reached before failing")
	t.Breakf("Test fails as FailNow()")
	t.NewStep("This step will never be reached after failing")
}

func (s *FailsDemoSuite) TestBrokenStatusNoMessageStep(t provider.T) {
	t.Title("This test broken without message in a step")
	t.Description(`
		This Test will be broken.
		No any message expected there`)

	t.Tags("fail", "broken")

	t.WithNewStep("This step will be broken", func(sCtx provider.StepCtx) {
		sCtx.Broken()
	})
	t.NewStep("This step will be reached after failing")
}

func (s *FailsDemoSuite) TestBrokenNowStatusNoMessageStep(t provider.T) {
	t.Title("This test broken immediately without message in a step")
	t.Description(`
		This Test will be broken.
		No any message expected there`)

	t.Tags("fail", "broken")

	t.WithNewStep("This step will be broken", func(sCtx provider.StepCtx) {
		sCtx.BrokenNow()
	})
	t.NewStep("This step will never be reached after failing")
}

func (s *FailsDemoSuite) TestBrokenNowStatusMessageStep(t provider.T) {
	t.Title("This test broken immediately with message in a step")
	t.Description(`
		This Test will be broken.
		Error message:
		Test fails as FailNow()`)

	t.Tags("fail", "broken")

	t.WithNewStep("This step will be broken", func(sCtx provider.StepCtx) {
		sCtx.Breakf("Test fails as FailNow()")
	})
	t.NewStep("This step will never be reached after failing")
}

func (s *FailsDemoSuite) TestBrokenNowStatusMessageInnerStep(t provider.T) {
	t.Title("This test broken immediately with message in an inner step")
	t.Description(`
		This Test will be broken.
		Error message:
		Test fails as FailNow()`)

	t.Tags("fail", "broken", "nesting")

	t.WithNewStep("This step will be broken", func(sCtx provider.StepCtx) {
		sCtx.WithNewStep("Inner step", func(sCtx provider.StepCtx) {
			sCtx.Breakf("Test fails as FailNow()")
		})
	})
	t.NewStep("This step will never be reached after failing")
}

func TestFails(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(FailsDemoSuite))
}
