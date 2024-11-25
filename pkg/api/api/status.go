package api

type Status struct {
	Code    string `gorm:"column:code" json:"code"`
	Message string `gorm:"column:message" json:"message"`
}

var (
	SUCCESS = Status{"00", "Operation completed successfully."}
	INVALID = Status{"89", "Invalid data"}
	ERROR   = Status{"99", "Internal server error."}
)
