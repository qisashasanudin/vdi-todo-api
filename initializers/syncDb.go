package initializers

import (
	"vdi/todo-api/entities/todo"
	"vdi/todo-api/entities/user"
)

func SyncDb() {
	// Migrate tables to database
	DB.AutoMigrate(&todo.Todo{})
	DB.AutoMigrate(&user.User{})
}
