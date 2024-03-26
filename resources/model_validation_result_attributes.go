package resources

type ValidationResultAttributes struct {
	Valid bool          `json:"valid"`
	Error *ErrorDetails `json:"error,omitempty"`
}

type ErrorDetails struct {
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
