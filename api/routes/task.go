package routes

import (
	"task/api/controller"
	"task/infrastructure"
)

type TaskRoute struct {
	Controller controller.TaskController
	Handler    infrastructure.GinRouter
}

func NewTaskRoute(
	controller controller.TaskController,
	handler infrastructure.GinRouter,
) TaskRoute {
	return TaskRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (p TaskRoute) Setup() {
	task := p.Handler.Gin.Group("/tasks")
	{
		task.POST("/", p.Controller.AddTask)
		task.GET("/:id", p.Controller.GetTask)
		task.DELETE("/:id", p.Controller.DeleteTask)
		task.PUT("/:id", p.Controller.UpdateTask)
	}
}
