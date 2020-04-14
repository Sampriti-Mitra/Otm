package dtos

import "encoding/json"

type CollabRequest struct {
	ProjectTitle  string            `json:"project_title" binding:"required"`
	Videolink     string            `json:"videolink"`
	CreatedBy     string            `json:"created_by"`
	Members       []string          `json:"members"`
	MembersStatus map[string]string `json:"members_status"`
	Model
}

type CollabResponse struct {
	ProjectTitle  string            `json:"project_title" binding:"required"`
	Videolink     string            `json:"videolink"`
	CreatedBy     string            `json:"created_by"`
	Members       []string          `json:"members"`
	MembersStatus map[string]string `json:"members_status"`
	Model
}
type Collab struct {
	ProjectTitle  string          `json:"project_title" binding:"required"`
	Videolink     string          `json:"videolink"`
	CreatedBy     string          `json:"created_by"`
	Members       json.RawMessage `json:"members"`
	MembersStatus json.RawMessage `json:"members_status"`
	Model
}

func (CollabRequest) TableName() string {
	return "collab"
}

func (Collab) TableName() string {
	return "collab"
}
