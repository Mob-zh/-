package repositories

import (
	"attendance_uniapp/initializer"
	"attendance_uniapp/models"
	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

// NewClassRepository 返回一个 ClassRepository 实例
func NewClassRepository() *ClassRepository {
	return &ClassRepository{DB: initializer.DB}
}

func (*ClassRepository) CreateClass(class *models.Class) error {
	return initializer.DB.Create(class).Error
}

func (*ClassRepository) GetClassById(classId string) (*models.Class, error) {
	class := &models.Class{}
	return class, initializer.DB.First(class, classId).Error
}

func (*ClassRepository) GetStudentListByClassId(classId string) ([]*models.Student, error) {
	var students []*models.Student
	return students, initializer.DB.Table("student_classes").Select("students.*").Joins("JOIN students ON student_classes.student_id = students.id").Where("student_classes.class_id = ?", classId).Find(&students).Error
}

func (*ClassRepository) GetClassListByStudentId(studentId string) ([]*models.Class, error) {
	var classes []*models.Class
	return classes, initializer.DB.Table("student_classes").Select("classes.*").Joins("JOIN classes ON student_classes.class_id = classes.id").Where("student_classes.student_id = ?", studentId).Find(&classes).Error
}

func (ClassRepo *ClassRepository) DeleteClassById(classId string) error {
	return initializer.DB.Where("class_id = ?", classId).Delete(&models.Class{}).Error

}
