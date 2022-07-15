package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	UserModel users.UserModel `json:"-"`
}

func NewUserModelValidator() users.Validators {
	userModelValidator := UserModelValidator{}
	return &userModelValidator
}

func (v *UserModelValidator) GetUsername() string {
	return v.User.Username
}

func (v *UserModelValidator) GetPassword() string {
	return v.User.Password
}

func (v *UserModelValidator) GetUserModel() *users.UserModel {
	return &v.UserModel
}

func (v *UserModelValidator) Bind(c *gin.Context) error {
	bind := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(v, bind); err != nil {
		return err
	}

	v.UserModel.Username = v.User.Username
	v.UserModel.Email = v.User.Email
	if err := v.UserModel.SetPassword(v.User.Password); err != nil {
		return err
	}
	return nil
}

type LoginValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	UserModel users.UserModel `json:"-"`
}

func NewLoginValidator() users.Validators {
	loginValidator := LoginValidator{}
	return &loginValidator
}

func (v *LoginValidator) GetUsername() string {
	return v.User.Username
}

func (v *LoginValidator) GetPassword() string {
	return v.User.Password
}

func (v *LoginValidator) GetUserModel() *users.UserModel {
	return &v.UserModel
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	bind := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(v, bind); err != nil {
		return err
	}

	v.UserModel.Username = v.User.Username
	return nil
}