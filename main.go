package main

import (
	"vdi/todo-api/entities/todo"
	"vdi/todo-api/entities/user"
	"vdi/todo-api/handlers"
	"vdi/todo-api/initializers"
	"vdi/todo-api/middleware"

	"github.com/gin-gonic/gin"
)

// Layers:
// - main
// - handler
// - service
// - repository
// - db (gorm)
// - mysql

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	// Initiate layers (repository, service & handler)
	userRepository := user.NewRepository(initializers.DB)
	userService := user.NewService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	todoRepository := todo.NewRepository(initializers.DB)
	todoService := todo.NewService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	// Initiate Gin
	router := gin.Default()
	v1 := router.Group("/v1")

	// Auth endpoints
	v1.POST("/auth/register", userHandler.CreateUser)
	v1.POST("/auth/login", userHandler.Login)
	v1.POST("/auth/logout", middleware.RequireAuth, userHandler.Logout)

	// User endpoints
	// v1.POST("/users", middleware.RequireAuth, userHandler.CreateUser)
	// v1.GET("/users", middleware.RequireAuth, userHandler.GetUsers)
	// v1.GET("/users/:id", middleware.RequireAuth, userHandler.GetUserById)
	// v1.PUT("/users/:id", middleware.RequireAuth, userHandler.UpdateUser)
	// v1.DELETE("/users/:id", middleware.RequireAuth, userHandler.DeleteUser)

	// Todo endpoints
	v1.POST("/todos", middleware.RequireAuth, todoHandler.CreateTodo)
	v1.GET("/todos", middleware.RequireAuth, todoHandler.GetTodos)
	v1.GET("/todos/:id", middleware.RequireAuth, todoHandler.GetTodoById)
	v1.PUT("/todos/:id", middleware.RequireAuth, todoHandler.UpdateTodo)
	v1.DELETE("/todos/:id", middleware.RequireAuth, todoHandler.DeleteTodo)

	// Run
	router.Run("localhost:8080")
}
