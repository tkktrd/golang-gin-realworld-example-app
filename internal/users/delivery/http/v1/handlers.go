package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
	"tkktrd/golang-gin-realworld-example-app/pkg/common"
)

type usersHandlers struct {
	useCase        users.UseCase
	userValidator  users.Validators
	loginValidator users.Validators
	userSerializer users.UserSerializer
}

func NewUsersHandlers(useCase users.UseCase, userValidator users.Validators, loginValidator users.Validators, userSerializer users.UserSerializer) users.Handlers {
	return &usersHandlers{
		useCase: useCase,
		userValidator: userValidator,
		loginValidator: loginValidator,
		userSerializer: userSerializer,
	}
}

func (h *usersHandlers) Registration() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userModelValidator := h.userValidator
		if err := userModelValidator.Bind(ctx); err != nil {
			e := common.NewStatusUnprocessableEntityError()
			ctx.JSON(e.Status(), e)
			return
		}

		err := h.useCase.Registration(ctx, userModelValidator.GetUserModel())
		if err != nil {
			e := common.ParseErrors(err)
			ctx.JSON(e.Status(), e)
			return
		}

		response, err := h.userSerializer.Response(userModelValidator.GetUserModel())
		if err != nil {
			e := common.NewStatusUnprocessableEntityError()
			ctx.JSON(e.Status(), e)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{ "user": response })
	}
}

func (h *usersHandlers) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginValidator := h.loginValidator
		if err := loginValidator.Bind(ctx); err != nil {
			e := common.NewStatusUnprocessableEntityError()
			ctx.JSON(e.Status(), e)
			return
		}

		response, err := h.userSerializer.Response(loginValidator.GetUserModel())
		if err != nil {
			e := common.NewStatusUnprocessableEntityError()
			ctx.JSON(e.Status(), e)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{ "user": response })
	}
}