package dtos

type ErrorResponse struct {
	Base
	Error Error `json:"error"`
}

type Error struct {
	ErrorCode    interface{} `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
}

type DeletedResponse struct {
	Id string `json:"id"`
	Base
}
