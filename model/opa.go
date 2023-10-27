package model

type OpaRequest struct {
	Input Input `json:"input"`
}

type OpaResponse struct {
	Allowed bool `json:"result"`
}

type Input struct {
	Path            []string `json:"path"`
	Method          string   `json:"method"`
	IsAuthenticated bool     `json:"authenticated"`
	Role            string   `json:"role"`
}
