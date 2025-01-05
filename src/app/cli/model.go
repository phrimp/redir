package cli

type Command struct {
	Action string `json:"action"`
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
}
