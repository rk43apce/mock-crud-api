package controllers

import (
	"mock-crud-api/utils"
	"net/http"

	// Adjust this import path based on your project structure
	"github.com/gin-gonic/gin"
)

// LoginHandler handles user login and generates JWT tokens
func LoginHandler(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON payload to creds struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Example: Hardcoded credentials (replace with your database validation)
	if creds.Username == "admin" && creds.Password == "password" {
		// Generate JWT token
		token, err := utils.GenerateJWT(creds.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// DashboardHandler handles the protected dashboard route
func DashboardHandler(c *gin.Context) {
	// user := c.MustGet("username").(string) // Extracted from the middleware

	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(400, gin.H{"message": "Cookie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the dashboard!",
		"user":    cookie,
	})
}
