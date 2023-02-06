package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]User, error)
	FindById(ID int) (User, error)
	FindByEmail(email string) (User, error)
	Create(userRequest CreateUserRequest) (User, error)
	Update(ID int, userRequest UpdateUserRequest) (User, error)
	Delete(ID int) (User, error)
	Register(registerRequest RegisterRequest) (User, string, error)
	Login(loginRequest LoginRequest) (User, string, error)
}

type service struct {
	repository Repository
}

func generateJwt(userData User) (string, error) {
	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userData.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	return s.repository.FindAll()

}

func (s *service) FindById(ID int) (User, error) {
	return s.repository.FindById(ID)
}

func (s *service) FindByEmail(email string) (User, error) {
	return s.repository.FindByEmail(email)
}

func (s *service) Create(b CreateUserRequest) (User, error) {

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(*b.Password), 10)
	if err != nil {
		return User{}, err
	}

	// Create a user
	user := User{
		Email:    *b.Email,
		Password: string(hash),
	}

	return s.repository.Create(user)
}

func (s *service) Update(ID int, b UpdateUserRequest) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	if b.Email != nil {
		user.Email = *b.Email
	}

	if b.Password != nil {
		user.Password = *b.Password
	}

	return s.repository.Update(user)
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	return s.repository.Delete(user)
}

func (s *service) Register(b RegisterRequest) (User, string, error) {

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(*b.Password), 10)
	if err != nil {
		return User{}, "", err
	}

	// Create a user
	user, err := s.repository.Create(User{
		Email:    *b.Email,
		Password: string(hash),
	})
	if err != nil {
		return user, "", err
	}

	// Generate a JWT
	tokenString, err := generateJwt(user)
	if err != nil {
		return user, "", err
	}

	// return user and token
	return user, tokenString, nil
}

func (s *service) Login(b LoginRequest) (User, string, error) {
	user, err := s.repository.FindByEmail(*b.Email)
	if err != nil {
		return user, "", err
	}

	// Compare the received password hash with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*b.Password))
	if err != nil {
		return user, "", err
	}

	// Generate a JWT
	tokenString, err := generateJwt(user)
	if err != nil {
		return user, "", err
	}

	// return user and token
	return user, tokenString, nil
}
