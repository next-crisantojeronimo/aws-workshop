package dto

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse HTTP Response template used for REST API responses
type HTTPResponse struct {
	Status  int
	Message string
	Content *[]TagResponse
}

// NewHTTPResponse Create a new HTTPResponse object
func NewHTTPResponse(status int, message string, content *[]TagResponse) *HTTPResponse {
	return &HTTPResponse{
		Status:  status,
		Message: message,
		Content: content,
	}
}

// SendResponse prepare and send http json response
func SendResponse(writer http.ResponseWriter, httpResponse *HTTPResponse) {
	responseJSON, _ := json.Marshal(httpResponse)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpResponse.Status)
	writer.Write(responseJSON)
}
