package initializer

import "attendance_uniapp/repositories"

var ClassRepository *repositories.ClassRepository
var CourseRepository *repositories.CourseRepository
var StudentRepository *repositories.StudentRepository
var TeacherRepository *repositories.TeacherRepository
var AttendanceRepository *repositories.AttendanceRepository

func initRepositories() {
	ClassRepository = repositories.NewClassRepository()
	CourseRepository = repositories.NewCourseRepository()
	StudentRepository = repositories.NewStudentRepository()
	TeacherRepository = repositories.NewTeacherRepository()
	AttendanceRepository = repositories.NewAttendanceRepository()
}
