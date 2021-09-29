package main

import (
	"task/api/controller"
	"task/api/repository"
	"task/api/routes"
	"task/api/service"
	"task/infrastructure"
	"task/models"
)

func main() {

	router := infrastructure.NewGinRouter()                     //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // databse has been initialized and configured
	taskRepository := repository.NewTaskRepository(db)          // repository are being setup
	taskService := service.NewTaskService(taskRepository)       // service are being setup
	taskController := controller.NewTaskController(taskService) // controller are being set up
	taskRoute := routes.NewTaskRoute(taskController, router)    // task routes are initialized
	taskRoute.Setup()                                           // task routes are being setup

	db.DB.AutoMigrate(&models.Task{}) // migrating Task model to datbase table
	router.Gin.Run(":8000")           //server started on 8000 port
}
