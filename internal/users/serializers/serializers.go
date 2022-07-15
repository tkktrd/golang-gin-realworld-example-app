package serializers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tkktrd/golang-gin-realworld-example-app/config"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
	"tkktrd/golang-gin-realworld-example-app/pkg/common"
)

type UserSerializer struct {
	config *config.Config
}

func NewUserSerializer(config *config.Config) users.UserSerializer {
	return &UserSerializer{ config: config }
}

func (s *UserSerializer) Response(ctx *gin.Context) (users.UserResponse, error) {
	model := ctx.MustGet("user_model").(*users.UserModel)

	res := users.UserResponse { Username: model.Username, Email: model.Email }
	token, err := common.GenerateJwtToken(model.ID, model.Email, s.config)
	if err != nil {
		return res, fmt.Errorf("failed to generate token, %v", err)
	}
	res.Token = token
	return res, nil
}