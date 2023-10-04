package models

// PythonAPIRequest represents the structure of the request being sent to the Python API.
type PythonAPIRequest struct {
	Text string `json:"text"`
}

// PythonAPIResponse encapsulates the response received from the Python API.
type PythonAPIResponse struct {
	Response string `json:"response"`
}

// RequestEnvelope is a standardized structure wrapping around a PythonAPIRequest.
type RequestEnvelope struct {
	Request PythonAPIRequest `json:"request"`
}

// ResponseEnvelope is a standardized structure wrapping around a PythonAPIResponse.
type ResponseEnvelope struct {
	Response PythonAPIResponse `json:"response"`
}

// APIError encapsulates standardized error attributes.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// APIResponse represents a generic API response structure, which can hold
// either a successful data payload or an error message.
type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *APIError   `json:"error,omitempty"`
}
