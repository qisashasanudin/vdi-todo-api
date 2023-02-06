package handlers

import (
	"net/http"
	"strconv"

	"vdi/todo-api/entities/user"
	"vdi/todo-api/utils"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUsers(c *gin.Context) {
	usersFromDB, err := h.userService.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(usersFromDB) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	var convertedUsers []user.UserResponse
	for _, b := range usersFromDB {
		convertedUser := utils.ConvertToUserResponse(b)
		convertedUsers = append(convertedUsers, convertedUser)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedUsers})
}

func (h *userHandler) GetUserById(c *gin.Context) {
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	userFromDB, err := h.userService.FindById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if userFromDB.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	convertedUser := utils.ConvertToUserResponse(userFromDB)
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedUser})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var newUserRequest user.CreateUserRequest
	err := c.ShouldBindJSON(&newUserRequest)
	utils.HandleValidationError(c, err)

	newUser, err := h.userService.Create(newUserRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newUser})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	var newUserRequest user.UpdateUserRequest
	err := c.ShouldBindJSON(&newUserRequest)
	utils.HandleValidationError(c, err)

	newUser, err := h.userService.Update(id, newUserRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newUser})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	userFromDB, err := h.userService.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	convertedUser := utils.ConvertToUserResponse(userFromDB)
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedUser})
}

func (h *userHandler) Register(c *gin.Context) {
	var registerRequest user.RegisterRequest
	err := c.ShouldBindJSON(&registerRequest)
	utils.HandleValidationError(c, err)

	newUser, token, err := h.userService.Register(registerRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the JWT as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 60*60*24, "", "", false, true)

	convertedUser := utils.ConvertToUserResponse(newUser)
	c.IndentedJSON(http.StatusCreated, gin.H{"data": convertedUser})
}

func (h *userHandler) Login(c *gin.Context) {
	var loginRequest user.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	utils.HandleValidationError(c, err)

	user, token, err := h.userService.Login(loginRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the JWT as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 60*60*24, "", "", false, true)

	convertedUser := utils.ConvertToUserResponse(user)
	c.IndentedJSON(http.StatusOK, gin.H{"data": convertedUser})
}

func (h *userHandler) Logout(c *gin.Context) {
	// Delete the JWT in the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Logout success"})
}
