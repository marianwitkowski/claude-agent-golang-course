package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	symbols = "!@#$%^&*()-_=+{}[]\\|;:'\",.<>/?"
)

type PasswordRequest struct {
	Length     int    `json:length`
	Complexity string `json:complexity`
}

type PasswordResponse struct {
	Password string `json:password`
}

func generatePassword(length int, complexity string) string {
	var validChars string
	switch complexity {
	case "simple":
		validChars = letters + digits
	case "complex":
		validChars = letters + digits + symbols
	default:
		validChars = letters + digits + symbols
	}
	password := make([]byte, length)
	for i := range password {
		password[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(password)

}

func passwordHandler(c *gin.Context) {
	var req PasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	password := generatePassword(req.Length, req.Complexity)

	resp := PasswordResponse{Password: password}
	c.JSON(200, resp)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		param1 := c.DefaultQuery("param1", "BRAK")
		c.JSON(200, gin.H{
			"answer": "Hello world!",
			"param1": param1,
		})
	})

	r.POST("/message", func(c *gin.Context) {
		txt := c.PostForm("txt")
		c.JSON(201, gin.H{
			"answer": txt,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("dane")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/password", passwordHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "nie ma tu nic"})
	})

	r.Run(":8000")

}
