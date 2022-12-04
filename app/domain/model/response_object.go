package model

type (
	ResponseData struct {
		IsError   bool        `json:"is_error"`
		ErrorCode string      `json:"error_code"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data,omitempty"`
	}

	ErrorItem struct {
		Code  string `json:"code"`
		Extra string `json:"extra"`
	}
)
