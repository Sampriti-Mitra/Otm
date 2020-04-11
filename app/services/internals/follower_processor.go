package internals

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	"otm/app/interfaces"
	"otm/app/repository"
	"strconv"
)

type followerProcessor struct {
}

func (l followerProcessor) Create(ctx *gin.Context, request dtos.FollowRequest) (dtos.FollowResponse, error) {
	var response dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	iError := repo.Create(ctx, &request)
	if iError != nil {
		return response, iError
	}

	response.Model = request.Model
	response.RequestBy = request.RequestBy
	response.RequestTo = request.RequestTo
	return response, nil
}

func (l followerProcessor) List(ctx *gin.Context, userId string) ([]dtos.FollowResponse, error) {
	var response []dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	err := repo.FindAll(ctx, &response, map[string]interface{}{
		"request_to": userId,
		"status":     "accept",
	})
	return response, err
}

func (l followerProcessor) ListFollowerRequest(ctx *gin.Context, userId string) ([]dtos.FollowResponse, error) {
	var response []dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	err := repo.FindAll(ctx, &response, map[string]interface{}{
		"request_to": userId,
		"status":     "pending",
	})
	return response, err
}

func (l followerProcessor) ListFollowing(ctx *gin.Context, userId string) ([]dtos.FollowResponse, error) {
	var response []dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	err := repo.FindAll(ctx, &response, map[string]interface{}{
		"request_by": userId,
		"status":     "accept",
	})
	return response, err
}

func (l followerProcessor) Get(ctx *gin.Context, userId string, requestBy string) (dtos.FollowResponse, error) {
	var response dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	err := repo.Find(ctx, &response, map[string]interface{}{
		"request_to": userId,
		"request_by": requestBy,
	})
	return response, err
}

func (l followerProcessor) UpdateAccept(ctx *gin.Context, profileId string, request_by string) (dtos.FollowResponse, error) {
	var response dtos.FollowResponse
	repo := repository.GetFollowerRepo()

	respGet, err := l.Get(ctx, profileId, request_by)
	if respGet.Status != "pending" {
		return response, errors.New("Can't accept request")
	}

	attributes := map[string]interface{}{
		"status": "accept",
	}
	_, err = repo.Transaction(
		ctx,
		func() (data interface{}, iError error) {

			err := repo.Update(ctx, &response, attributes, map[string]interface{}{
				"request_to": profileId,
				"request_by": request_by,
			})
			if err != nil {
				return response, err
			}

			// followed musician
			resp, iError := GetLoginProcessor().GetByUsername(ctx, profileId)
			if iError != nil {
				return
			}
			aboutResp, iError := GetAboutProcessor().Get(ctx, resp.Id)
			if iError != nil {
				return
			}
			var aboutReq dtos.AboutRequest
			aboutReq.Followers = aboutResp.Followers + 1
			aboutReq.Following = aboutResp.Following
			updatedResp, iError := GetAboutProcessor().UpdateFollows(ctx, aboutReq, resp.Id)
			if iError != nil {
				return
			}

			// the one following the musician
			resp, iError = GetLoginProcessor().GetByUsername(ctx, request_by)
			if iError != nil {
				return
			}
			aboutResp, iError = GetAboutProcessor().Get(ctx, resp.Id)
			if iError != nil {
				return
			}
			aboutReq = dtos.AboutRequest{}
			aboutReq.Followers = aboutResp.Followers
			aboutReq.Following = aboutResp.Following + 1
			updatedResp, iError = GetAboutProcessor().UpdateFollows(ctx, aboutReq, resp.Id)
			if iError != nil {
				return
			}

			fmt.Println(updatedResp)
			return data, iError
		},
	)
	if err != nil {
		return response, err
	}

	return response, err
}

func (l followerProcessor) UpdateReject(ctx *gin.Context, profileId string, request_by string) (dtos.FollowResponse, error) {
	var response dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	respGet, err := l.Get(ctx, profileId, request_by)
	if respGet.Status != "pending" && respGet.Status != "accept" {
		return response, errors.New("Can't reject request")
	}

	attributes := map[string]interface{}{
		"status": "reject",
	}

	if respGet.Status == "pending" {
		err = repo.Update(ctx, &response, attributes, map[string]interface{}{
			"request_to": profileId,
			"request_by": request_by,
		})
		if err != nil {
			return response, err
		}
	} else {
		_, err = repo.Transaction(
			ctx,
			func() (data interface{}, iError error) {

				err := repo.Update(ctx, &response, attributes, map[string]interface{}{
					"request_to": profileId,
					"request_by": request_by,
				})
				if err != nil {
					return response, err
				}

				// followed musician
				resp, iError := GetLoginProcessor().GetByUsername(ctx, profileId)
				if iError != nil {
					return
				}
				aboutResp, iError := GetAboutProcessor().Get(ctx, resp.Id)
				if iError != nil {
					return
				}
				var aboutReq dtos.AboutRequest
				aboutReq.Followers = aboutResp.Followers - 1
				aboutReq.Following = aboutResp.Following
				updatedResp, iError := GetAboutProcessor().UpdateFollows(ctx, aboutReq, resp.Id)
				if iError != nil {
					return
				}

				// the one following the musician
				resp, iError = GetLoginProcessor().GetByUsername(ctx, request_by)
				if iError != nil {
					return
				}
				aboutResp, iError = GetAboutProcessor().Get(ctx, resp.Id)
				if iError != nil {
					return
				}
				aboutReq = dtos.AboutRequest{}
				aboutReq.Followers = aboutResp.Followers
				aboutReq.Following = aboutResp.Following - 1
				updatedResp, iError = GetAboutProcessor().UpdateFollows(ctx, aboutReq, resp.Id)
				if iError != nil {
					return
				}

				fmt.Println(updatedResp)
				return data, iError
			},
		)
		if err != nil {
			return response, err
		}
	}

	return response, err
}

func (l followerProcessor) Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse {
	var errorResp dtos.DeletedResponse
	var users dtos.Users
	repo := repository.GetLoginRepo()
	_, err := repo.Transaction(
		ctx,
		func() (data interface{}, iError error) {

			// Create func for namespace
			errResp := repo.Delete(ctx, &users, map[string]interface{}{
				"id": requestId,
			})
			if errResp != nil {
				return nil, errResp
			}

			//Create Domain Model
			userId, iError := strconv.Atoi(requestId)
			if iError != nil {
				return
			}
			iError = GetAboutProcessor().Delete(ctx, userId)
			if iError != nil {
				return
			}

			return data, iError
		},
	)
	if err != nil {
		errorResp.Success = false
	} else {
		errorResp.Success = true
	}
	errorResp.Id = fmt.Sprintf("%d", users.Id)
	errorResp.Error = err
	return errorResp
}

func GetFollowerProcessor() interfaces.IFollower {
	return followerProcessor{}
}
