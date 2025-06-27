//go:build examples_new
// +build examples_new

package suite_demo

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ParametrizedTestParallelDemo struct {
	suite.Suite
}

func (s *ParametrizedTestParallelDemo) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Parametrized-Parallel")
}

func (s *ParametrizedTestParallelDemo) TestParameterizedParallel(t provider.T) {
	t.Title("Parent Test-Parallel")
	t.Description(`
		This test is parent for all Parametrized-Parallel`)

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized-Parallel 1#%d", newI), func(t provider.T) {
			t.Epic("Demo")
			t.Feature("Parametrized-Parallel")
			t.Description(fmt.Sprintf("`This test checks that 1 Equal %d`", newI))
			t.Tag("Parametrized-Parallel")
			t.Parallel()
			t.WithNewStep(fmt.Sprintf("Step %d", i), func(ctx provider.StepCtx) {
				ctx.Require().Equal(1, newI)
			})
		})
	}

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized-Parall	el 2#%d", newI), func(t provider.T) {
			t.Epic("Demo")
			t.Feature("Parametrized-Parallel")
			t.Description(fmt.Sprintf("`This test checks that 1 Less %d`", newI))
			t.Tag("Parametrized-Parallel")
			t.Parallel()
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func(ctx provider.StepCtx) {
				ctx.Require().Less(1, newI)
			})
		})
	}
}

func TestParametrizedParallelDemo(t *testing.T) {
	suite.RunSuite(t, new(ParametrizedTestParallelDemo))
}

type ParametrizedTestDemo struct {
	suite.Suite
}

func (s *ParametrizedTestDemo) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Parametrized")
}

func (s *ParametrizedTestDemo) TestParameterized(t provider.T) {
	t.Title("Parent Test")
	t.Description(`
		This test is parent for all Parametrized`)

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized 1#%d", newI), func(t provider.T) {
			t.Epic("Demo")
			t.Feature("Parametrized")
			t.Description(fmt.Sprintf("`This test checks that 1 Equal %d`", newI))
			t.Tag("Parametrized")
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func(ctx provider.StepCtx) {
				ctx.Require().Equal(1, newI)
			})
		})
	}

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized 2#%d", newI), func(t provider.T) {
			t.Epic("Demo")
			t.Feature("Parametrized")
			t.Description(fmt.Sprintf("`This test checks that 1 Less %d`", newI))
			t.Tag("Parametrized")
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func(ctx provider.StepCtx) {
				if newI == 4 {
					panic("WHOOPS")
				}
				ctx.Require().Less(1, newI)
			})
		})
	}
}

func TestParametrizedDemo(t *testing.T) {
	suite.RunSuite(t, new(ParametrizedTestDemo))
}
