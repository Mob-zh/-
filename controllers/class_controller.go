package controllers

import (
	"attendance_uniapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TeacherCreateClass(ctx *gin.Context) {
	var input struct {
		ClassName  string `json:"class_name" binding:"required"`
		TeacherId  string `json:"teacher_id" binding:"required"`
		CourseId   string `json:"course_id" binding:"required"`
		CourseTime string `json:"course_time" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建班级的业务逻辑
	newClass := models.Class{
		ClassName: input.ClassName,
		TeacherID: input.TeacherID,
	}

	if err := models.CreateClass(&newClass); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Class created successfully", "class": newClass})
}
