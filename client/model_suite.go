package client

type Suite struct {
	Id            int64  `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	Preconditions string `json:"preconditions,omitempty"`
	Position      int32  `json:"position,omitempty"`
	CasesCount    int32  `json:"cases_count,omitempty"`
	ParentId      int64  `json:"parent_id,omitempty"`
	Created       string `json:"created,omitempty"`
	Updated       string `json:"updated,omitempty"`
}
