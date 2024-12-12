package routes

import (
	"attendance_uniapp/controllers"
	"attendance_uniapp/initializer"
	"attendance_uniapp/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	studentCtrler := controllers.NewStudentController(initializer.StudentService, initializer.ClassService)
	teacherCtrler := controllers.NewTeacherController(initializer.TeacherService, initializer.ClassService, initializer.CourseService)
	commonCtrler := controllers.NewCommonController(initializer.StudentService, initializer.TeacherService)

	r := gin.Default()
	student := r.Group("/student/")
	{
		//学生登录
		student.POST("/login", commonCtrler.LoginHandler("student"))
		//以下需验证JWT
		student.Use(middleware.ValidateJWT(), middleware.ValidateRole("student"))
		{

			//学生主页
			student.GET("/home", studentCtrler.StudentGetHomeHandler)
			//学生修改密码
			student.PATCH("/changePwd", commonCtrler.ChangePwdHandler)

			class := student.Group("/:class_id")
			{
				class.Use(middleware.ValidateClassExist(), middleware.ValidateUserClassMatched("student"))
				//学生加入班级
				class.POST("/join", studentCtrler.StudentJoinClassHandler)
				//学生退出班级
				class.POST("/quit", studentCtrler.StudentQuitFromClassHandler)
				//学生获取班级信息
				class.GET("/info", studentCtrler.StudentGetClassInfoHandler)
				//学生在该班级中进行签到操作
				class.POST("/sign")
			}

		}

	}
	teacher := r.Group("/teacher/")
	{
		//老师登录
		teacher.POST("/login", commonCtrler.LoginHandler("teacher"))
		teacher.Use(middleware.ValidateJWT(), middleware.ValidateRole("teacher"))
		{
			//老师主页
			teacher.GET("/home", teacherCtrler.TeacherGetHome)
			//老师修改密码
			teacher.PATCH("/changePwd", commonCtrler.ChangePwdHandler)
			class := teacher.Group("/:class_id")
			{
				class.Use(middleware.ValidateClassExist(), middleware.ValidateUserClassMatched("teacher"))
				//老师创建班级
				class.POST("/create", teacherCtrler.TeacherCreateClassHandler)
				//老师删除班级
				class.DELETE("/delete", teacherCtrler.TeacherDeleteClassHandler)
				//老师获取某一班级信息
				class.GET("/info", teacherCtrler.TeacherGetClassInfoHandler)
				//老师在该班级中进行考勤操作
				class.POST("/sign")
				{
					class.POST("/sign/saveRecord")
					class.POST("/sign/deleteRecord")

				}
				//老师在该班级中获取考勤记录
				class.GET("/fetchRecord")
				//老师手动补签(待定)
				class.POST("/signRecord/fix")
			}
		}
	}
	return r
}
