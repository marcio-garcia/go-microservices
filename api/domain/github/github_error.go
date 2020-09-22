package github

// ErrorResponse - error response model
type ErrorResponse struct {
	StatusCode       int            `json:"status_code"`
	Message          string         `json:"message"`
	DocumentationURL string         `json:"documentation_url"`
	Errors           []ErrorDetails `json:"errors"`
}

func (r ErrorResponse) Error() string {
	return r.Message
}

// ErrorDetails - error details
type ErrorDetails struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
