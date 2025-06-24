//go:build examples_new
// +build examples_new

package suite_demo

import (
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type JSONStruct struct {
	Message string `json:"message"`
}

type AttachmentTestDemoSuite struct {
	suite.Suite
}

func (s *AttachmentTestDemoSuite) TestAttachment(t provider.T) {
	t.Epic("Demo")
	t.Feature("Attachments")
	t.Title("Test Attachments")
	t.Description(`
		The test body and all its steps can contain attachments`)

	t.Tags("Attachments", "Steps")

	attachmentText := `THIS IS A TEXT ATTACHMENT`
	t.WithAttachments(allure.NewAttachment("Text Attachment", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Step A")
	var ExampleJson = JSONStruct{"this is JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.WithAttachments(allure.NewAttachment("Json Attachment for Step A", allure.JSON, attachmentJSON))
	t.Step(step)
}

type AttachmentDemoSuite struct {
	suite.Suite
}

func (s *AttachmentDemoSuite) BeforeAll(t provider.T) {
	// this action will create a step in Set up
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	t.WithAttachments(allure.NewAttachment("Text Attachment for Before Suite", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Before Suite Step")
	var ExampleJson = JSONStruct{"This is BeforeAll JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.WithAttachments(allure.NewAttachment("Json Attachment for Before Suite Step", allure.JSON, attachmentJSON))
	t.Step(step)
}

func (s *AttachmentDemoSuite) BeforeEach(t provider.T) {
	// this action will create a step in Set up
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	t.WithAttachments(allure.NewAttachment("Text Attachment for Before Test", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Before Test Step")
	var ExampleJson = JSONStruct{"This is BeforeEach JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.WithAttachments(allure.NewAttachment("Json Attachment for Before Test Step", allure.JSON, attachmentJSON))
	t.Step(step)
}

func (s *AttachmentDemoSuite) AfterAll(t provider.T) {
	// this action will create a step in Tear down
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	t.WithAttachments(allure.NewAttachment("Text Attachment for After Suite", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("After Suite Step")
	var ExampleJson = JSONStruct{"This is AfterAll JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.WithAttachments(allure.NewAttachment("Json Attachment for After Suite Step", allure.JSON, attachmentJSON))
	t.Step(step)
}

func (s *AttachmentDemoSuite) AfterEach(t provider.T) {
	// this action will create a step in Tear down
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	t.WithAttachments(allure.NewAttachment("Text Attachment for After Test", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("After Test Step")
	var ExampleJson = JSONStruct{"This is AfterEach JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.WithAttachments(allure.NewAttachment("Json Attachment for After Test Step", allure.JSON, attachmentJSON))
	t.Step(step)
}

func (s *AttachmentDemoSuite) TestBeforeAfterAttachments(t provider.T) {
	t.Epic("Demo")
	t.Feature("Attachments")
	t.Title("Test with Before/After Attachments")
	t.Description(`
		Test "Set up", Test "Tear down", suite "Set up", suite "Tear down" 
		and all steps inside can contain attachments`)

	t.Tags("Attachments", "BeforeAfter", "Steps")
}

type NestedAttachmentDemoSuite struct {
	suite.Suite
}

func (s *NestedAttachmentDemoSuite) BeforeAll(t provider.T) {
	t.WithNewStep("Before Suite Step", func(ctx provider.StepCtx) {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		ctx.WithAttachments(allure.NewAttachment("Text Attachment for Before Suite Step", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) AfterAll(t provider.T) {
	t.WithNewStep("After Suite Step", func(ctx provider.StepCtx) {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		ctx.WithAttachments(allure.NewAttachment("Text Attachment for After Suite Step", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) BeforeEach(t provider.T) {
	t.WithNewStep("Before Test Step", func(ctx provider.StepCtx) {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		ctx.WithAttachments(allure.NewAttachment("Text Attachment for Before Test Step", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) AfterEach(t provider.T) {
	t.WithNewStep("After Test Step", func(ctx provider.StepCtx) {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		ctx.WithAttachments(allure.NewAttachment("Text Attachment for After Test Step", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) TestNestedAttachment(t provider.T) {
	t.Epic("Demo")
	t.Feature("Attachments")
	t.Title("Test NestedAttachments")
	t.Description(`
		Suite "Set up", Test "Set up", Test "Tear down", Suite "Tear down" have steps with attachments.`)

	t.Tags("Attachments", "Nesting", "Steps", "BeforeAfter")

	t.WithNewStep("TestNestedAttachment step", func(ctx provider.StepCtx) {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		ctx.WithAttachments(allure.NewAttachment("Text Attachment for TestNestedAttachment step", allure.Text, []byte(attachmentText)))
	})
}

func TestAttachments(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(AttachmentTestDemoSuite))
	suite.RunSuite(t, new(AttachmentDemoSuite))
	suite.RunSuite(t, new(NestedAttachmentDemoSuite))
}
