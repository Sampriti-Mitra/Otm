package dtos

type FollowRequest struct {
	RequestBy string `json:"request_by" binding:"required"`
	RequestTo string `json:"request_to"`
	Model
}

type FollowResponse struct {
	RequestBy string `json:"request_by"`
	RequestTo string `json:"request_to"`
	Status    string `json:"status"`
	Model
}

func (FollowRequest) TableName() string {
	return "follower_request"
}

func (FollowResponse) TableName() string {
	return "follower_request"
}
