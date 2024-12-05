package controllers

import (
	"attendance_uniapp/models"
	"attendance_uniapp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TeacherCreateClass(ctx *gin.Context) {
	var input struct {
		ClassName string `json:"class_name" binding:"required"`
		CourseId  string `json:"course_id" binding:"required"`
		ClassTime string `json:"course_time" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 构造新增数据
	classId, err := utils.GenerateClassId(input.ClassName, input.ClassTime, ctx.GetString("user_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成班级ID失败"})
	}

	queryForCourse, err := models.GetCourseById(input.CourseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取课程信息失败"})
		return
	}
	courseName := queryForCourse.CourseName

	queryForTeacher, err := models.GetTeacherById(ctx.GetString("user_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取教师信息失败"})
		return
	}
	teacherName := queryForTeacher.TeacherName

	// 创建班级
	classToCreate := models.Class{
		ClassId:            classId,
		ClassName:          input.ClassName,
		TeacherId:          ctx.GetString("user_id"),
		CourseName:         courseName,
		CourseId:           input.CourseId,
		ClassTime:          input.ClassTime,
		TeacherName:        teacherName,
		IsClassChecking:    false,
		CheckingEndTime:    time.Now(),
		TotalCheckingTimes: 0,
	}

	if err := models.CreateClass(&classToCreate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Class created successfully"})
}
