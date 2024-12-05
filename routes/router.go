package routes

import (
	"attendance_uniapp/controllers"
	"attendance_uniapp/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	student := r.Group("/student/")
	{
		//学生登录
		student.POST("/login", controllers.Login("student"))
		//以下需验证JWT
		student.Use(middleware.ValidateJWT())
		{
			//学生修改密码
			student.POST("/changePwd", controllers.ChangePwd)
			{
				class := student.Group("/<string:classId>")
				{
					//学生加入班级
					class.POST("/join")
					//学生退出班级
					class.POST("/quit")
					//学生获取班级信息
					class.GET("/info")
					//学生在该班级中进行签到操作
					class.POST("/sign")
				}
			}

		}

	}
	teacher := r.Group("/teacher/")
	{
		//老师登录
		teacher.POST("/login", controllers.Login("teacher"))
		teacher.Use(middleware.ValidateJWT())
		{
			//老师修改密码
			teacher.POST("/changePwd", controllers.ChangePwd)
			class := teacher.Group("/<string:classId>")
			{
				//老师创建班级
				class.POST("/create")
				//老师删除班级
				class.POST("/delete")
				//老师获取班级信息
				class.GET("/info")
				//老师在该班级中进行考勤操作
				class.POST("/sign")
				//老师在该班级中获取考勤记录
				class.GET("/signRecord")
				//老师手动补签
				class.POST("/signRecord/fix")

			}
		}
	}
	return r
}
