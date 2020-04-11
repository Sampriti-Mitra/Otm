package dtos

type AboutRequest struct {
	About     string `json:"about" binding:"required"`
	UserId    int    `json:"user_id"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

type AboutResponse struct {
	About     string `json:"about"`
	UserId    int    `json:"user_id"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Success   bool   `json:"success"`
	Model
}
type About struct {
	About     string `json:"about"`
	UserId    int    `json:"user_id"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Model
}

func (AboutRequest) TableName() string {
	return "about"
}

func (About) TableName() string {
	return "about"
}
