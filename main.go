package main

import (
	"log"
	"os"
	"task/api/controller"
	"task/api/repository"
	"task/api/routes"
	"task/api/service"
	"task/infrastructure"
	"task/models"
)

var (
	db     infrastructure.Database
	router infrastructure.GinRouter
)

func init() {
	infrastructure.LoadEnv()
	router = infrastructure.NewGinRouter()
	db = infrastructure.NewDatabase()
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := controller.NewTaskController(taskService)
	taskRoute := routes.NewTaskRoute(taskController, router)
	taskRoute.Setup()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	err := db.DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("Task didn't migrate to db")
	}

	if err := router.Gin.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
