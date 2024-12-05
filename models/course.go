package models

import "attendance_uniapp/initializer"

type Course struct {
	CourseId   string `json:"course_id" gorm:"primaryKey"` //课程id
	CourseName string `json:"course_name"`                 //课程名称
}

func GetCourseById(courseId string) (*Course, error) {
	course := &Course{}
	// 查询课程，如果没有找到，返回错误
	if err := initializer.DB.Where("course_id = ?", courseId).First(&course).Error; err != nil {
		return nil, err
	}
	// 返回查询到的课程以及 nil 错误
	return course, nil
}
