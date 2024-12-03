package dto

type Request struct {
	RequestId   string `json:"request_id" binding:"required"`
	RequestTime string `json:"request_time" binding:"required"`
	PageNumber  int    `json:"page_number"`
	PageSize    int    `json:"page_size"`
}
