package middleware

import (
	"fmt"
	"github.com/D3FiX4M/go-restfull/initializers"
	"github.com/D3FiX4M/go-restfull/models"
	"github.com/D3FiX4M/go-restfull/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"status": "fail", "message": "You are not logged in"})
		}

		config, _ := initializers.LoadConfig(".")

		sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"status": "fail", "message": err.Error()})
			return
		}
		var user models.User

		result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
