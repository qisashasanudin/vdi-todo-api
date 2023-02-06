package todo

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Todo, error)
	FindById(ID int) (Todo, error)
	Create(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(todo Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *repository) FindById(ID int) (Todo, error) {
	var todo Todo
	err := r.db.Find(&todo, ID).Error
	return todo, err
}

func (r *repository) Create(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *repository) Update(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error
	return todo, err
}

func (r *repository) Delete(todo Todo) (Todo, error) {
	err := r.db.Delete(&todo).Error
	return todo, err
}
