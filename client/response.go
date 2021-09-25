package client

import (
	"net/http"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	Operation      string `json:"operation,omitempty"`
	RequestURL     string `json:"url,omitempty"`
	Method         string `json:"method,omitempty"`
	Payload        []byte `json:"-"`
}

func NewAPIResponse(r *http.Response) *APIResponse {
	response := &APIResponse{Response: r}
	return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {
	response := &APIResponse{Message: errorMessage}
	return response
}
