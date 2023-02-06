package utils

import (
	"vdi/todo-api/entities/todo"
	"vdi/todo-api/entities/user"
)

func ConvertToUserResponse(b user.User) user.UserResponse {
	return user.UserResponse{
		ID:       b.ID,
		Email:    b.Email,
		Password: b.Password,
	}
}

func ConvertToTodoResponse(b todo.Todo) todo.TodoResponse {
	return todo.TodoResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Completed:   b.Completed,
	}
}
