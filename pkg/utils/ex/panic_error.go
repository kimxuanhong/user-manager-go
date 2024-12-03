package ex

import "encoding/json"

type panicError struct {
	code    string
	message string
}

func New(code string, message string) error {
	return &panicError{code, message}
}

func (e *panicError) Error() string {
	return e.code + ": " + e.message
}

func (e *panicError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e *panicError) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e)
}
