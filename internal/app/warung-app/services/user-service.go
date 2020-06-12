package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/afadhitya/warung-backend/internal/app/warung-app/config"
	"github.com/afadhitya/warung-backend/internal/app/warung-app/util"
	"net/http"

	"github.com/afadhitya/warung-backend/internal/app/warung-app/models"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	address := c.PostForm("address")
	phoneNumber := c.PostForm("phoneNumber")
	email := c.PostForm("email")

	if isUsernameNotAvailabel(&username) {
		util.SetContextError(c, http.StatusBadRequest, "username not available")
		return
	}

	doPasswordEncription(&password)

	user := models.User{
		Username:    username,
		Password:    password,
		Name:        name,
		Address:     address,
		PhoneNumber: phoneNumber,
		Email:       email,
	}

	config.DB.Save(&user)

	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "User has saved to DB",
		"resourceId": user.ID,
	})
}

func SignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	doPasswordEncription(&password)

	user := getOneUser(&username)

	if isUserNotAvailable(&user) {
		util.SetContextError(c, http.StatusNotFound, "User Not Found")
		return
	}

	if user.Password != password {
		util.SetContextError(c, http.StatusBadRequest, "wrong password")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login Success",
	})
}

func getOneUser(username *string) models.User {
	var user models.User
	config.DB.Where("username = ?", (*username)).Find(&user)
	return user
}

func doPasswordEncription(password *string) {
	var sha = sha1.New()
	sha.Write([]byte(*password))
	var encrypted = sha.Sum(nil)
	(*password) = fmt.Sprintf("%x", encrypted)
}

func isUsernameNotAvailabel(username *string) bool {
	var count int
	var user models.User
	config.DB.Where("username = ?", (*username)).Find(&user).Count(&count)
	return (count > 0)
}

func isUserNotAvailable(user *models.User) bool{
	return (user.ID == 0)
}

