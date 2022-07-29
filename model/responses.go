package model

// APIResponse is standard response for all APIs.
type APIResponse struct {
	StatusCode int
	Data       interface{}
}

type APIResponseWithError struct {
	StatusCode int
	Data       interface{}
	Error      error
}

// ErrorStatus ...
type ErrorStatus struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

// ErrorBody ...
type ErrorBody struct {
	Status ErrorStatus `json:"status,omitempty"`
}

// ResponseResult ...
type ResponseResult struct {
	StatusCode int         `json:"status_code,omitempty"`
	Data       interface{} `json:"data"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
}
