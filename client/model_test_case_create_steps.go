package client

type TestCaseCreateSteps struct {
	Action         string   `json:"action,omitempty"`
	ExpectedResult string   `json:"expected_result,omitempty"`
	Data           string   `json:"data,omitempty"`
	Position       int32    `json:"position,omitempty"`
	Attachments    []string `json:"attachments,omitempty"`
}
