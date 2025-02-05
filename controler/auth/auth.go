package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lekchan000/go-jwt-api/orm"
	"golang.org/x/crypto/bcrypt"
)

// RegisterBody represents the incoming JSON request
type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody

	// Bind incoming JSON to RegisterBody struct
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure DB is initialized
	if orm.Db == nil {
		log.Println("❌ Database connection is nil in Register function")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}

	// Check if user exists
	var userExist orm.User
	result := orm.Db.Where("Username = ?", json.Username).First(&userExist)
	if result.Error == nil { // User exists
		c.JSON(http.StatusOK, gin.H{"Status": "error", "message": "User Exists"})
		return
	}

	// Encrypt password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	if err != nil {
		log.Println("❌ Error encrypting password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	// Create new user
	user := orm.User{
		Username: json.Username,
		Password: string(encryptedPassword),
		Fullname: json.Fullname,
		Avatar:   json.Avatar,
	}

	// Insert into database
	createResult := orm.Db.Create(&user)
	if createResult.Error != nil {
		log.Println("❌ Error creating user:", createResult.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "error", "message": "User creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "OK", "message": "User Create Success", "userID": user.ID})
}
