package repositories

import (
	"attendance_uniapp/global"
	"attendance_uniapp/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

// NewCourseRepository 返回一个 CourseRepository 实例
func NewCourseRepository() *CourseRepository {
	return &CourseRepository{global.DB}
}

func (*CourseRepository) GetCourseById(courseId string) (*models.Course, error) {
	course := &models.Course{}
	return course, global.DB.Where("course_id = ?", courseId).First(&course).Error
}
