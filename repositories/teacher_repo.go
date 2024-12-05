package repositories

import (
	"attendance_uniapp/initializer"
	"attendance_uniapp/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	DB *gorm.DB
}

// NewTeacherRepository 返回一个 TeacherRepository 实例
func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{DB: initializer.DB}
}

func (*TeacherRepository) GetClassesByTeacherId(teacherId string) ([]models.Class, error) {
	var classes []models.Class
	// 查询teacherId对应的classes
	return classes, initializer.DB.Where("teacher_id = ?", teacherId).Find(&classes).Error
}

func (*TeacherRepository) GetTeacherById(teacherId string) (*models.Teacher, error) {
	var teacher models.Teacher
	return &teacher, initializer.DB.Where("teacher_id = ?", teacherId).First(&teacher).Error
}
