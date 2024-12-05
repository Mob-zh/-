package controllers

import (
	"attendance_uniapp/initializer"
	"attendance_uniapp/models"
	"attendance_uniapp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// 设置通用查询条件
var userModel interface{}
var passwordField string

func Login(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// input结构体 获取请求参数
		var input struct {
			UserId string `json:"user_id" binding:"required"`
			Pwd    string `json:"pwd" binding:"required"`
		}

		// 检查响应数据json是否正确
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 根据角色选择模型和密码字段
		if role == "student" {
			userModel = &models.Student{}
			passwordField = "StudentPwd"
		} else if role == "teacher" {
			userModel = &models.Teacher{}
			passwordField = "TeacherPwd"
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
			return
		}

		// 查询用户
		err := initializer.DB.Where(role+"_id = ?", input.UserId).First(userModel).Error
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid " + role + " ID or password"})
			return
		}

		// 获取用户模型的密码字段，使用反射
		userValue := reflect.ValueOf(userModel).Elem()
		passwordHash := userValue.FieldByName(passwordField).String()

		// 验证密码
		if !utils.CheckPassword(input.Pwd, passwordHash) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid " + role + " ID or password"})
			return
		}

		// 生成 token
		token, err := utils.GenerateJWT(input.UserId, role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// 响应数据
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func ChangePwd(ctx *gin.Context) {
	// input结构体 获取请求参数
	var input struct {
		OldPwd         string `json:"old_pwd"`
		NewPwd         string `json:"new_pwd"`
		RepeatedNewPwd string `json:"repeated_new_pwd"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证新密码的两次输入是否相同和新密码是否与旧密码相同
	if input.NewPwd != input.RepeatedNewPwd || input.NewPwd == input.OldPwd {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "New passwords do not match or are the same as the old password"})
	}

	// 根据上下文取参
	role := ctx.GetString("role")
	userId := ctx.GetString("user_id")

	// 根据角色选择模型和密码字段
	if role == "student" {
		userModel = &models.Student{}
		passwordField = "StudentPwd"
	} else if role == "teacher" {
		userModel = &models.Teacher{}
		passwordField = "TeacherPwd"
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// 查询用户
	if err := initializer.DB.Where(role+"_id = ?", userId).First(userModel).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
		return
	}

	// 获取用户模型的密码字段，使用反射
	userValue := reflect.ValueOf(userModel).Elem()
	passwordHash := userValue.FieldByName(passwordField).String()

	// 验证密码
	if !utils.CheckPassword(input.OldPwd, passwordHash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid " + role + " ID or password"})
		return
	}
	// 更新密码
	toUpdateHashedPwd, _ := utils.HashPassword(input.NewPwd)
	if err := initializer.DB.Model(userModel).Update(passwordField, toUpdateHashedPwd).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}
	// 响应数据
	ctx.JSON(http.StatusOK, gin.H{"msg": "Password updated successfully"})
}
