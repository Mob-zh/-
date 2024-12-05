package middleware

import (
	"attendance_uniapp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ValidateJWT JWT验证中间件
func ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求中的 Authorization header
		tokenString := ctx.GetHeader("Authorization")
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			ctx.Abort()
			return
		}

		// 解析 JWT
		userId, role, err := utils.ParseJWT(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}
		// 将 userId 、role 放到上下文中
		ctx.Set("user_id", userId)
		ctx.Set("role", role)

		ctx.Next()
	}
}

func ValidateRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleFromContext := ctx.GetString("role")
		if roleFromContext != role {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Insufficient permissions"})
			ctx.Abort()
			return

		}
		ctx.Next()

	}
}
