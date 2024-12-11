package repositories

import (
	"attendance_uniapp/global"
	"attendance_uniapp/models"
	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

// NewClassRepository 返回一个 ClassRepository 实例
func NewClassRepository() *ClassRepository {
	return &ClassRepository{DB: global.DB}
}

func (*ClassRepository) CreateClass(class *models.Class) error {
	return global.DB.Create(class).Error
}

func (*ClassRepository) GetClassById(classId string) (*models.Class, error) {
	class := &models.Class{}
	return class, global.DB.First(class, classId).Error
}

func (*ClassRepository) GetStudentListByClassId(classId string) ([]*models.Student, error) {
	var students []*models.Student
	return students, global.DB.Table("student_classes").Select("students.*").Joins("JOIN students ON student_classes.student_id = students.id").Where("student_classes.class_id = ?", classId).Find(&students).Error
}

func (*ClassRepository) GetClassListByStudentId(studentId string) ([]*models.Class, error) {
	var classes []*models.Class
	return classes, global.DB.Table("student_classes").Select("classes.*").Joins("JOIN classes ON student_classes.class_id = classes.id").Where("student_classes.student_id = ?", studentId).Find(&classes).Error
}

func (ClassRepo *ClassRepository) DeleteClassById(classId string) error {
	return global.DB.Where("class_id = ?", classId).Delete(&models.Class{}).Error

}
