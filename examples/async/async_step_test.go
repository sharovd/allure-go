//go:build async
// +build async

package async

import (
	"fmt"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type StepAsyncDemo struct {
	suite.Suite
}

func (s *StepAsyncDemo) BeforeEach(t provider.T) {
	t.Epic("Async")
	t.Feature("Async Steps")
	t.Tags("async", "suite", "steps")
}

func (s *StepAsyncDemo) TestAsyncStepsDemo(t provider.T) {
	t.Title("Test with async steps")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.Logf("Step 2 Stopped At: %s", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepsWithInnerStepsDemo(t provider.T) {
	t.Title("Test with async steps and inner steps")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepsPanicInAStepDemo(t provider.T) {
	t.Title("Test with async steps and panic in a step")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		panic("Whoops")
	})
}

func (s *StepAsyncDemo) TestAsyncStepsPanicInAnInnerStepDemo(t provider.T) {
	t.Title("Test with async steps and panic in an inner step")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			panic("Whoops")
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepsFailedInAStepDemo(t provider.T) {
	t.Title("Test with async steps and fail by assert in a step")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		ctx.Assert().False(true)
	})
}

func (s *StepAsyncDemo) TestAsyncStepsFailedInAnInnerStepDemo(t provider.T) {
	t.Title("Test with async steps and fail by assert in an inner step")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			ctx.Assert().False(true)
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

type AsyncSuiteStepDemo struct {
	suite.Suite
}

func (s *AsyncSuiteStepDemo) BeforeEach(t provider.T) {
	t.Epic("Async")
	t.Feature("Async Steps")
	t.Tags("async", "suite", "steps")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepsParallelDemo(t provider.T) {
	t.Title("Test with async steps (parallel)")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.Logf("Step 2 Stopped At: %s", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepsWithInnerStepsParallelDemo(t provider.T) {
	t.Title("Test with async steps and inner steps (parallel)")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepsPanicInAStepParallelDemo(t provider.T) {
	t.Title("Test with async steps and panic in a step (parallel)")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		panic("Whoops")
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepsPanicInAnInnerStepParallelDemo(t provider.T) {
	t.Title("Test with async steps and panic in an inner step (parallel)")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))

		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			panic("Whoops")
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepsFailedInAStepParallelDemo(t provider.T) {
	t.Title("Test with async steps and fail by assert in a step (parallel)")

	t.Parallel()
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		ctx.Assert().False(true)
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepsFailedInAnInnerStepParallelDemo(t provider.T) {
	t.Title("Test with async steps and fail by assert in an inner step (parallel)")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			ctx.Assert().False(true)
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func TestStepAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(StepAsyncDemo))
}

func TestSuiteStepAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(AsyncSuiteStepDemo))
}
