//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type StepDemoSuite struct {
	suite.Suite
}

func (s *StepDemoSuite) TestAddSteps(t provider.T) {
	t.Epic("Demo")
	t.Feature("Steps")
	t.Title("Add steps to Allure report")
	t.Description(`
		Step A, Step B and Step C will be add to Allure report
		Step B contains paramB with value
		Step C contains paramC with value`)

	t.Tags("Steps")

	stepA := allure.NewSimpleStep("Step A")
	t.Step(stepA)

	stepB := allure.NewStep("Step B", // Step's Name
		allure.Passed,                           // Step Status
		allure.GetNow(),                         // Step Start
		allure.GetNow(),                         // Step Finish
		allure.NewParameters("paramB", "value")) // Step Parameters

	t.Step(stepB)

	stepC := allure.NewSimpleStep("Step C")
	stepC.Start = allure.GetNow()
	stepC.WithNewParameters("paramC", "value")
	stepC.Stop = allure.GetNow()
	t.Step(stepC)
}

func (s *StepDemoSuite) TestAddStepsWithDifferentStatuses(t provider.T) {
	t.Epic("Demo")
	t.Feature("Steps")
	t.Title("Add steps with different statuses to Allure report")
	t.Description(`
		Step A (Passed), Step B (Failed), Step C (Skipped) and Step D (Passed) will be add to Allure report`)

	t.Tags("Steps")

	t.Step(allure.NewSimpleStep("Step A").Passed())
	t.Step(allure.NewSimpleStep("Step B").Failed())
	t.Step(allure.NewSimpleStep("Step C").Skipped())

	stepD := allure.NewSimpleStep("Step D").Begin()
	time.Sleep(1 * time.Second) // Do some
	stepD = stepD.Finish().Passed()
	t.Step(stepD)
}

func (s *StepDemoSuite) TestInnerStep(t provider.T) {
	t.Epic("Demo")
	t.Layer("Layer")
	t.Feature("Steps")
	t.Title("Add inner steps to existed steps in Allure report")
	t.Description(`
		Step A is parent step for Step B and Step C
		Step D is parent step for Step E and Step F
		Step F contains paramF with value
		Call order will be saved in allure report
		A -> (B, C), D -> (E, F)`)

	t.Tags("Steps", "Nesting")

	// use allure.NewSimpleStep constructor
	stepA := allure.NewSimpleStep("Step A")
	stepB := allure.NewSimpleStep("Step B")
	stepC := allure.NewSimpleStep("Step C")
	stepA.WithChild(stepB)
	stepA.WithChild(stepC)
	t.Step(stepA)

	// use InnerStep function
	stepD := allure.NewSimpleStep("Step D")
	t.Step(stepD)
	stepD.WithChild(allure.NewSimpleStep("Step E"))
	stepF := allure.NewStep("Step F", // Step's Name
		allure.Passed,                           // Step Status
		allure.GetNow(),                         // Step Start
		allure.GetNow(),                         // Step Finish
		allure.NewParameters("paramF", "value")) // Step Parameters
	stepF.WithParent(stepD)

	// forward way
	stepG := allure.NewSimpleStep("Step G")
	stepH := allure.NewSimpleStep("Step H")
	stepI := allure.NewSimpleStep("Step I")
	stepG.WithChild(stepH)
	stepG.WithChild(stepI)
}

func TestStepDemo(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(StepDemoSuite))
}
