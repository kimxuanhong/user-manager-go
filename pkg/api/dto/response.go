package dto

type Response struct {
	ResponseId   string      `json:"response_id"`
	ResponseTime string      `json:"response_time"`
	ResponseCode Status      `json:"response_code"`
	Data         interface{} `json:"data"`
}
