package initializer

import "attendance_uniapp/services"

var StudentService *services.StudentService
var CourseService *services.CourseService
var TeacherService *services.TeacherService
var ClassService *services.ClassService
var AttendanceService *services.AttendanceService

func initServices() {
	StudentService = services.NewStudentService(StudentRepository, ClassRepository, AttendanceRepository)
	CourseService = services.NewCourseService(CourseRepository)
	TeacherService = services.NewTeacherService(TeacherRepository)
	ClassService = services.NewClassService(ClassRepository, StudentRepository, StudentClassRepository)
	AttendanceService = services.NewAttendanceService(AttendanceRepository, ClassRepository)
}
