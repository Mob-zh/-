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
	ClassServ      *services.ClassService
	TeacherServ    *services.TeacherService
	CourseServ     *services.CourseService
	AttendanceServ *services.AttendanceService
}

func NewTeacherController(TeacherServ *services.TeacherService, ClassServ *services.ClassService, CourseServ *services.CourseService, AttendanceServ *services.AttendanceService) *TeacherController {
	return &TeacherController{ClassServ: ClassServ, TeacherServ: TeacherServ, CourseServ: CourseServ, AttendanceServ: AttendanceServ}
}

func (TeacherCtrl *TeacherController) TeacherGetHomeHandler(ctx *gin.Context) {
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
		CheckingEndTime:    time.Now(),
		TotalCheckingTimes: 0,
	}
	for err := TeacherCtrl.ClassServ.ClassRepo.CreateClass(&classToCreate); err != nil; classToCreate.ClassId, _ = utils.GenerateClassId(input.ClassName, input.ClassTime, ctx.GetString("user_id")) {
		err = TeacherCtrl.ClassServ.ClassRepo.CreateClass(&classToCreate)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Class created successfully"})
}

// TeacherGetClassInfoHandler 老师查看班级详情
func (TeacherCtrl *TeacherController) TeacherGetClassInfoHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	queryForClass, _ := TeacherCtrl.ClassServ.GetClassByIdService(classId)
	ctx.JSON(http.StatusOK, gin.H{
		"class_id":    queryForClass.ClassId,
		"class_name":  queryForClass.ClassName,
		"course_name": queryForClass.CourseName,
	})
}

// TeacherGetClassStudentInfoHandler 老师查看该班级所有学生信息
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

// TeacherDeleteClassHandler 教师删除班级
func (TeacherCtrl *TeacherController) TeacherDeleteClassHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	if err := TeacherCtrl.ClassServ.DeleteClassByIdService(classId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// TeacherStartToCheckHandler 教师开始考勤
func (TeacherCtrl *TeacherController) TeacherStartToCheckHandler(ctx *gin.Context) {
	var input struct {
		CheckingSeconds int `json:"checking_seconds"` // 考勤时长
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请设置考勤时长"})
		return
	}

	//初始化redis相应数据并获取签到码
	signInCode, err := TeacherCtrl.AttendanceServ.StartToCheckAndGetSignInCodeService(ctx.Param("class_id"), input.CheckingSeconds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
	}
	//响应签到码
	ctx.JSON(http.StatusOK, gin.H{"sign_in_code": signInCode})
}

// TeacherGetSignedInCountHandler 获取实时签到计数情况，计数包括班级总人数，以“{{已签到人数}}/{{班级总人数}}”的形式返回
func (TeacherCtrl *TeacherController) TeacherGetSignedInCountHandler(ctx *gin.Context) {
	signInCount, err := TeacherCtrl.AttendanceServ.GetSignedInCountService(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
	}
	ctx.JSON(http.StatusOK, gin.H{"signed_in_count": signInCount})
}

// TeacherGetSignInSummaryHandler 教师获取班级当前签到汇总
func (TeacherCtrl *TeacherController) TeacherGetSignInSummaryHandler(ctx *gin.Context) {
	//直接返回一个组件
	component, err := TeacherCtrl.AttendanceServ.GetSignInSummaryService(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
	}

	ctx.JSON(http.StatusOK, gin.H{"component": component})

}
