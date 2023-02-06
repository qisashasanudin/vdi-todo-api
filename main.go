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

	// Auth endpoints
	router.POST("/auth/register", userHandler.CreateUser)
	router.POST("/auth/login", userHandler.Login)
	router.POST("/auth/logout", middleware.RequireAuth, userHandler.Logout)

	// User endpoints
	// router.POST("/users", middleware.RequireAuth, userHandler.CreateUser)
	// router.GET("/users", middleware.RequireAuth, userHandler.GetUsers)
	// router.GET("/users/:id", middleware.RequireAuth, userHandler.GetUserById)
	// router.PUT("/users/:id", middleware.RequireAuth, userHandler.UpdateUser)
	// router.DELETE("/users/:id", middleware.RequireAuth, userHandler.DeleteUser)

	// Todo endpoints
	router.POST("/todos", middleware.RequireAuth, todoHandler.CreateTodo)
	router.GET("/todos", middleware.RequireAuth, todoHandler.GetTodos)
	router.GET("/todos/:id", middleware.RequireAuth, todoHandler.GetTodoById)
	router.PUT("/todos/:id", middleware.RequireAuth, todoHandler.UpdateTodo)
	router.DELETE("/todos/:id", middleware.RequireAuth, todoHandler.DeleteTodo)

	// Run
	router.Run("localhost:8080")
}
