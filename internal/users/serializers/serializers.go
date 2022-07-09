package serializers

import "tkktrd/golang-gin-realworld-example-app/internal/users"

type UserSerializer struct {

}

func NewUserSerializer() users.UserSerializer {
	return &UserSerializer{}
}

func (s *UserSerializer) Response(userModel *users.UserModel) users.UserResponse {
	res := users.UserResponse{
		Username: userModel.Username,
		Email: userModel.Email,
		Token: "",
	}
	return res
}