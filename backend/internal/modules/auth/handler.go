package auth

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	_ = c.ShouldBindJSON(&req)

	// TEMP: hardcoded users
	if req.Email == "admin@foodos.com" {
		token, _ := GenerateToken("1", "ADMIN")
		c.JSON(200, gin.H{"token": token, "role": "ADMIN"})
		return
	}

	token, _ := GenerateToken("2", "USER")
	c.JSON(200, gin.H{"token": token, "role": "USER"})
}
