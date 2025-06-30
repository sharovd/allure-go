package async

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MixedAsyncSuite struct {
	suite.Suite
}

func setInfoForAllureReport(t provider.T) {
	t.Epic("Async")
	t.Feature("Mixed Suite")
	t.Tags("async", "suite", "steps")
}

func (s *MixedAsyncSuite) TestMixedAsyncSuiteDemo1(t provider.T) {
	t.SkipOnPrint()
	testCases := []struct {
		testName string
	}{{"Test 1.1"}, {"Test 1.2"}, {"Test 1.3"}}

	for _, tc := range testCases {
		t.Run(tc.testName+" - Passed", func(t provider.T) {
			setInfoForAllureReport(t)
			name := tc.testName
			t.Parallel()
			t.NewStep(name + " - Step")
		})
	}
}

func (s *MixedAsyncSuite) TestMixedAsyncSuiteDemo2(t provider.T) {
	t.SkipOnPrint()
	t.Parallel()
	testCases := []struct {
		testName string
	}{{"Test 2.1"}, {"Test 2.2"}, {"Test 2.3"}}

	for _, tc := range testCases {
		t.Run(tc.testName+" - Passed", func(t provider.T) {
			setInfoForAllureReport(t)
			name := tc.testName
			t.Parallel()
			t.NewStep(name + " - Step")
		})
	}
}

func (s *MixedAsyncSuite) TestMixedAsyncSuiteDemo3(t provider.T) {
	t.SkipOnPrint()
	testCases := []struct {
		testName string
	}{{"Test 3.1"}, {"Test 3.2"}, {"Test 3.3"}}

	for _, tc := range testCases {
		t.Run(tc.testName+" - Failed", func(t provider.T) {
			setInfoForAllureReport(t)
			name := tc.testName
			t.Parallel()
			t.NewStep(name + " - Step")
			t.Fatalf("WHOOPS")
		})
	}
}

func (s *MixedAsyncSuite) TestMixedAsyncSuiteDemo4(t provider.T) {
	t.SkipOnPrint()
	testCases := []struct {
		testName string
	}{{"Test 4.1"}, {"Test 4.2"}, {"Test 4.3"}}

	for _, tc := range testCases {
		t.Run(tc.testName+" - Broken", func(t provider.T) {
			setInfoForAllureReport(t)
			name := tc.testName
			t.Parallel()
			t.NewStep(name + " - Step")
			panic("WHOOPS")
		})
	}
}

func TestMixedAsyncSuite(t *testing.T) {
	suite.RunSuite(t, new(MixedAsyncSuite))
}
