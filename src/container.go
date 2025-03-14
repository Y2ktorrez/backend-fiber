package src

import (
	"github.com/torrez/config"
	"github.com/torrez/src/handlers"
	"github.com/torrez/src/repository"
	"github.com/torrez/src/services"
)

type Container struct {
	UserRepo    *repository.UserRepository
	UserService *services.UserService
	UserHandler *handlers.UserHandler
}

func SetupContainer() *Container {

	//Users
	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	return &Container{
		//Users
		UserRepo:    userRepo,
		UserService: userService,
		UserHandler: userHandler,
	}
}
