package client

type TestCaseUpdateSteps struct {
	Action         string `json:"action,omitempty"`
	ExpectedResult string `json:"expected_result,omitempty"`
	Data           string `json:"data,omitempty"`
	Position       int32  `json:"position,omitempty"`
}
