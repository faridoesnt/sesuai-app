package ahttp

type Message struct {
	Str string `json:"message"`
}

// Success represents base responses structure if a request is success
type Success struct {
	Result   interface{}       `json:"data"`
	Metadata interface{}       `json:"_metadata,omitempty"`
	Header   map[string]string `json:"-"`
}

func OK() *Success {
	return &Success{
		Result: Message{
			Str: "OK",
		},
	}
}

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Debug   *ErrorDebug `json:"_debug,omitempty"`
}

type ErrorDebug struct {
	Trace   string `json:"trace,omitempty"`
	Message string `json:"err_message,omitempty"`
}
