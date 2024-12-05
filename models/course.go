package models

type Course struct {
	CourseId   string `json:"course_id" gorm:"primaryKey"` //课程id
	CourseName string `json:"course_name"`                 //课程名称
}
