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
	userSerializer users.UserSerializer
}

func NewUsersHandlers(useCase users.UseCase, userValidator users.Validators, userSerializer users.UserSerializer) users.Handlers {
	return &usersHandlers{ useCase: useCase, userValidator: userValidator, userSerializer: userSerializer }
}

func (h *usersHandlers) Registration() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userModelValidator := h.userValidator
		if err := userModelValidator.Bind(ctx); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, common.NewStatusUnprocessableEntityError())
			return
		}

		err := h.useCase.Registration(ctx, userModelValidator.GetUserModel())
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, common.NewStatusUnprocessableEntityError())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": h.userSerializer.Response(userModelValidator.GetUserModel()),
		})
	}
}