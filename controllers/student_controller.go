package controllers

import (
	"attendance_uniapp/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentController struct {
	StudentServ *services.StudentService
	ClassServ   *services.ClassService
}

func NewStudentController(StudentServ *services.StudentService, ClassServ *services.ClassService) *StudentController {
	return &StudentController{StudentServ: StudentServ, ClassServ: ClassServ}
}

// StudentGetClassInfoHandler 学生进入某一班级详情页
func (StudentCtrl *StudentController) StudentGetClassInfoHandler(ctx *gin.Context) {
	classId := ctx.Param("class_id")
	queryForClass, err := StudentCtrl.ClassServ.GetClassByIdService(classId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "请检查您的网络并稍后再试",
		})
	}
	if queryForClass != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"class_name":   queryForClass.ClassName,
			"course_name":  queryForClass.CourseName,
			"teacher_name": queryForClass.TeacherName,
			"class_time":   queryForClass.ClassTime,
			"classroom":    queryForClass.Classroom,
		})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "未找到该班级",
	})
}

// StudentGetHomeHandler 学生首页
func (StudentCtrl *StudentController) StudentGetHomeHandler(ctx *gin.Context) {
	classList, err := StudentCtrl.ClassServ.GetClassListByStudentIdService(ctx.GetString("user_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "请检查您的网络并稍后再试",
		})
	}
	classNameList := make([]string, 0)
	for _, class := range classList {
		classNameList = append(classNameList, class.ClassName)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"classes": classNameList,
	})
}

// StudentJoinClassHandler 学生加入班级
func (StudentCtrl *StudentController) StudentJoinClassHandler(ctx *gin.Context) {
	var input struct {
		ClassId string `json:"class_id"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// 调用服务层加入班级
	err := StudentCtrl.StudentServ.EnrollStudentInClassService(ctx.GetString("user_id"), input.ClassId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Class joined successfully"})
}

// StudentQuitFromClassHandler 学生退出班级
func (StudentCtrl *StudentController) StudentQuitFromClassHandler(ctx *gin.Context) {
	var input struct {
		ClassId string `json:"class_id"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// 调用服务层退出班级
	err := StudentCtrl.StudentServ.StudentQuitFromClassService(ctx.GetString("user_id"), input.ClassId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "请检查您的网络并稍后再试"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Class quit successfully"})
}

// StudentSigninHandler 学生签到操作
func (StudentCtrl *StudentController) StudentSigninHandler(ctx *gin.Context) {
	var input struct {
		SigninTime string `json:"signin_time"`
		ClassId    string `json:"class_id"`
	}

}
