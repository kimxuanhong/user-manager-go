package dto

type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	SUCCESS   = Status{"00", "Operation completed successfully."}
	NOT_FOUND = Status{"89", "Resource not found."}
	ERROR     = Status{"99", "Internal server error."}
)
