package routes

import (
	"task/api/controller"
	"task/infrastructure"
)

//TaskRoute -> Route for question module
type TaskRoute struct {
	Controller controller.TaskController
	Handler    infrastructure.GinRouter
}

//NewTaskRoute -> initializes new choice routes
func NewTaskRoute(
	controller controller.TaskController,
	handler infrastructure.GinRouter,
) TaskRoute {
	return TaskRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p TaskRoute) Setup() {
	task := p.Handler.Gin.Group("/tasks") //Router group
	{
		task.POST("/", p.Controller.AddTask)
		task.GET("/:id", p.Controller.GetTask)
		task.DELETE("/:id", p.Controller.DeleteTask)
		task.PUT("/:id", p.Controller.UpdateTask)
	}
}
