package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tkktrd/golang-gin-realworld-example-app/config"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
	"tkktrd/golang-gin-realworld-example-app/internal/users/delivery/http/v1"
	"tkktrd/golang-gin-realworld-example-app/internal/users/repository"
	"tkktrd/golang-gin-realworld-example-app/internal/users/serializers"
	"tkktrd/golang-gin-realworld-example-app/internal/users/usecase"
	"tkktrd/golang-gin-realworld-example-app/internal/users/validators"
)

type App struct {
	gin *gin.Engine
	config *config.Config
	db *gorm.DB
}

func NewApp(config *config.Config, db *gorm.DB) *App {
	app := App{
		gin: gin.Default(),
		config: config,
		db: db,
	}
	return &app
}

func (app *App) Run() {
	api := app.gin.Group("/api/v1")
	{
		v1.MapNewRoutes(api, app.createUserHandler())
	}
	app.gin.Run(":3000")
}

func (app *App) createUserHandler() users.Handlers {
	repository := repository.NewUsersRepository(app.db)
	return v1.NewUsersHandlers(
		usecase.NewUserUseCase(repository),
		validators.NewUserModelValidator(),
		validators.NewLoginValidator(),
		serializers.NewUserSerializer(app.config))
}