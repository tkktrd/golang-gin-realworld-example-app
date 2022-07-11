package users

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Token    string  `json:"token"`
}

type UserSerializer interface {
	Response(userModel *UserModel) (UserResponse, error)
}