package controller

import (
	"net/http"
	"strconv"
	"task/api/service"
	"task/models"
	"task/util"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service service.TaskService
}

func NewTaskController(s service.TaskService) TaskController {
	return TaskController{
		service: s,
	}
}

func (p *TaskController) AddTask(ctx *gin.Context) {
	var task models.Task
	ctx.ShouldBindJSON(&task)

	if task.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	if task.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
		return
	}
	if task.Author == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Author is required")
		return
	}
	err := p.service.Save(task)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create task")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Task")
}

func (p *TaskController) GetTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var task models.Task
	task.ID = id
	foundTask, err := p.service.Find(task)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Task")
		return
	}
	response := foundTask.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Task",
		Data:    &response})

}

func (p *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Task")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

func (p TaskController) UpdateTask(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var task models.Task
	task.ID = id

	taskRecord, err := p.service.Find(task)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Task with given id not found")
		return
	}
	ctx.ShouldBindJSON(&taskRecord)

	if taskRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	if taskRecord.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
		return
	}
	if taskRecord.Author == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Author is required")
		return
	}

	if err := p.service.Update(taskRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Task")
		return
	}
	response := taskRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Task",
		Data:    response,
	})
}
