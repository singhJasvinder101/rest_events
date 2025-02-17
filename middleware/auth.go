package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhJasvinder101/utils"
)

func Authenticate(c *gin.Context){
	token := c.GetHeader("Authorization")

    if token == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "message": "Unauthorized",
        })
        return
    }

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	
	c.Set("userId", userId)
	c.Next()
}

