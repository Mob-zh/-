package middleware

import (
	"attendance_uniapp/initializer"
	"attendance_uniapp/models"
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

func ValidateClassExist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		classId := ctx.Param("class_id")
		if err := initializer.DB.Where("class_id = ?", classId).First(&models.Class{}).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func ValidateUserClassMatched(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var isMatched bool
		classId := ctx.Param("class_id")
		isMatched = initializer.DB.Model(&models.Class{}).Select("1").Where(role+"_id = ? AND class_id = ?", ctx.GetString("user_id"), classId).Find(&isMatched).RowsAffected > 0
		if !isMatched {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "您没有权限访问该班级"})
			ctx.Abort()
		}
		ctx.Next()
	}

}
