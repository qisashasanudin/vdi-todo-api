package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindById(ID int) (User, error)
	FindByEmail(email string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindById(ID int) (User, error) {
	var user User
	err := r.db.Find(&user, ID).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) Delete(user User) (User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
