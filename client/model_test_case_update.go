package client

type TestCaseUpdate struct {
	Description    string                `json:"description,omitempty" yaml:"description,omitempty"`
	Preconditions  string                `json:"preconditions,omitempty" yaml:"preconditions,omitempty"`
	Postconditions string                `json:"postconditions,omitempty" yaml:"postconditions,omitempty"`
	Title          string                `json:"title,omitempty" yaml:"title,omitempty"`
	Severity       int32                 `json:"severity,omitempty" yaml:"severity,omitempty"`
	Priority       int32                 `json:"priority,omitempty" yaml:"priority,omitempty"`
	Behavior       int32                 `json:"behavior,omitempty" yaml:"behavior,omitempty"`
	Type_          int32                 `json:"type,omitempty" yaml:"type,omitempty"`
	Layer          int32                 `json:"layer,omitempty" yaml:"layer,omitempty"`
	IsFlaky        int32                 `json:"is_flaky,omitempty" yaml:"is_flaky,omitempty"`
	SuiteId        int64                 `json:"suite_id,omitempty" yaml:"suite_id,omitempty"`
	MilestoneId    int64                 `json:"milestone_id,omitempty" yaml:"milestone_id,omitempty"`
	Automation     int32                 `json:"automation,omitempty" yaml:"automation,omitempty"`
	Status         int32                 `json:"status,omitempty" yaml:"status,omitempty"`
	Params         []Params              `json:"params,omitempty" yaml:"params,omitempty"`
	Steps          []TestCaseUpdateSteps `json:"steps,omitempty" yaml:"steps,omitempty"`
	// A map of custom fields values (id => value)
	CustomField map[string]string `json:"custom_field,omitempty" yaml:"custom_field,omitempty"`
}
