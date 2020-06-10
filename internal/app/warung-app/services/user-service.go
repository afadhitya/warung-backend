package services

import (
	"crypto/sha1"
	"fmt"
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

	doPasswordEncription(&password)

	user := models.User{
		Username:    username,
		Password:    password,
		Name:        name,
		Address:     address,
		PhoneNumber: phoneNumber,
		Email:       email,
	}

	models.DB.Save(&user)

	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "User has saved to DB",
		"resourceId": user.ID,
	})
}

func doPasswordEncription(password *string) {
	var sha = sha1.New()
	sha.Write([]byte(*password))
	var encrypted = sha.Sum(nil)
	(*password) = fmt.Sprintf("%x", encrypted)
}
