package dtos

import "encoding/json"

type UploadRequest struct {
	Videolink string `json:"videolink"`
	Title     string `json:"title" binding:"required"`
	CreatedBy string `json:"created_by"`
	Tags      string `json:"tags"`
	Model
}

type UploadResponse struct {
	Videolink     string          `json:"videolink" binding:"required"`
	Title         string          `json:"title" binding:"required"`
	CreatedBy     string          `json:"created_by" binding:"required"`
	Tags          string          `json:"tags" binding:"required"`
	Applause      int             `json:"applause" binding:"required"`
	CountStreamed int             `json:"count_streamed" binding:"required"`
	Success       bool            `json:"success"`
	ApplaudedBy   json.RawMessage `json:"applauded_by"`
	Model
}

type Profile struct {
	AboutResponse
	Posts []UploadResponse
}

type Upload struct {
	RequestedBy string `json:"requested_by" binding:"required"`
	Model
}

type Applaud struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Model
}

func (UploadRequest) TableName() string {
	return "profile"
}

func (UploadResponse) TableName() string {
	return "profile"
}

func (Upload) TableName() string {
	return "profile"
}

func (Applaud) TableName() string {
	return "profile"
}
