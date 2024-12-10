package controllers

import (
	"attendance_uniapp/models"
	"attendance_uniapp/services"
	"attendance_uniapp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TeacherController struct {
	ClassServ   *services.ClassService
	TeacherServ *services.TeacherService
	CourseServ  *services.CourseService
}

func NewTeacherController(TeacherServ *services.TeacherService, ClassServ *services.ClassService, CourseServ *services.CourseService) *TeacherController {
	return &TeacherController{ClassServ: ClassServ, TeacherServ: TeacherServ, CourseServ: CourseServ}
}

func (TeacherCtrl *TeacherController) TeacherGetHome(ctx *gin.Context) {
	teacherId := ctx.Param("user_id")
	var classes []models.Class
	classes, err := TeacherCtrl.TeacherServ.TeacherRepo.GetClassesByTeacherId(teacherId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"classes": classes})
}
func (TeacherCtrl *TeacherController) TeacherCreateClassHandler(ctx *gin.Context) {
	var input struct {
		ClassName string `json:"class_name" binding:"required"`
		CourseId  string `json:"course_id" binding:"required"`
		ClassTime string `json:"course_time" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	classId, _ := utils.GenerateClassId(input.ClassName, input.ClassTime, ctx.GetString("user_id"))

	queryForCourse, err := TeacherCtrl.CourseServ.GetCourseByIdService(input.CourseId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "您的课程号有误"})
		return
	}
	courseName := queryForCourse.CourseName

	queryForTeacher, err := TeacherCtrl.TeacherServ.GetTeacherByIdService(ctx.GetString("user_id"))
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

	if err := TeacherCtrl.ClassServ.ClassRepo.CreateClass(&classToCreate); err != nil {
		TeacherCtrl.TeacherCreateClassHandler(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Class created successfully"})
}

// 老师查看班级详情
func (TeacherCtrl *TeacherController) TeacherGetClassInfoHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	queryForClass, _ := TeacherCtrl.ClassServ.GetClassByIdService(classId)
	ctx.JSON(http.StatusOK, gin.H{
		"class_id":    queryForClass.ClassId,
		"class_name":  queryForClass.ClassName,
		"course_name": queryForClass.CourseName,
	})
}

// 老师查看该班级所有学生信息
func (TeacherCtrl *TeacherController) TeacherGetClassStudentInfoHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	studentList, _ := TeacherCtrl.ClassServ.GetStudentListByClassIdService(classId)
	studentListIds := make([]string, 0)
	studentListNames := make([]string, 0)
	for _, student := range studentList {
		studentListIds = append(studentListIds, student.StudentId)
		studentListNames = append(studentListNames, student.StudentName)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"student_id":   studentListIds,
		"student_name": studentListNames,
	})
}

// 教师发起考勤
func (TeacherCtrl *TeacherController) TeacherStartAttendanceHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	var input struct {
		CheckingEndTime string `json:"checking_end_time"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "错误的请求数据"})
	}

}

// 教师删除班级
func (TeacherCtrl *TeacherController) TeacherDeleteClassHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	if err := TeacherCtrl.ClassServ.DeleteClassByIdService(classId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
