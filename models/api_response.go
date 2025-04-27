package models

type ApiResponsePattern struct {
	Status    string               `json:"status"`
	Data      interface{}          `json:"data"`
	Message   string               `json:"message"`
	Code      int                  `json:"code"`
	Meta      *ResponsePatternMeta `json:"meta,omitempty"`
	ErrorCode string               `json:"error_code,omitempty"`
}

type ResponsePatternMeta struct {
	TotalPages  int `json:"total_pages"`
	TotalItems  int `json:"total_items"`
	CurrentPage int `json:"current_page"`
}
