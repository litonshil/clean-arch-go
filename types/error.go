package types

type ValidationError struct {
	Error  string   `json:"validationError,omitempty"`
	Errors []string `json:"validationErrors,omitempty"`
}
