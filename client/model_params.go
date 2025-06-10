package client

type Params struct {
	Title  string   `json:"title,omitempty" yaml:"title,omitempty"`
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
