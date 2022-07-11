package serializers

import (
	"fmt"
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

func (s *UserSerializer) Response(userModel *users.UserModel) (users.UserResponse, error) {
	res := users.UserResponse { Username: userModel.Username, Email: userModel.Email }
	token, err := common.GenerateJwtToken(userModel.ID, userModel.Email, s.config)
	if err != nil {
		return res, fmt.Errorf("failed to generate token, %v", err)
	}
	res.Token = token
	return res, nil
}