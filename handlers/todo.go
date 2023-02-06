package handlers

import (
	"net/http"
	"strconv"

	"vdi/todo-api/entities/todo"
	"vdi/todo-api/utils"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService todo.Service
}

func NewTodoHandler(todoService todo.Service) *todoHandler {
	return &todoHandler{todoService}
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	todosFromDB, err := h.todoService.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(todosFromDB) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todos not found"})
		return
	}

	var convertedTodos []todo.TodoResponse
	for _, b := range todosFromDB {
		convertedTodo := utils.ConvertToTodoResponse(b)
		convertedTodos = append(convertedTodos, convertedTodo)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedTodos})
}

func (h *todoHandler) GetTodoById(c *gin.Context) {
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	todoFromDB, err := h.todoService.FindById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if todoFromDB.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	convertedTodo := utils.ConvertToTodoResponse(todoFromDB)
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedTodo})
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var newTodoRequest todo.CreateTodoRequest
	err := c.ShouldBindJSON(&newTodoRequest)
	utils.HandleValidationError(c, err)

	newTodo, err := h.todoService.Create(newTodoRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newTodo})
}

func (h *todoHandler) UpdateTodo(c *gin.Context) {
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	var newTodoRequest todo.UpdateTodoRequest
	err := c.ShouldBindJSON(&newTodoRequest)
	utils.HandleValidationError(c, err)

	newTodo, err := h.todoService.Update(id, newTodoRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newTodo})
}

func (h *todoHandler) DeleteTodo(c *gin.Context) {
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	todoFromDB, err := h.todoService.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	convertedTodo := utils.ConvertToTodoResponse(todoFromDB)
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedTodo})
}
