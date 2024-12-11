package repositories

import (
	"attendance_uniapp/global"
	"attendance_uniapp/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

// NewStudentRepository 返回一个 StudentRepository 实例
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{DB: global.DB}
}

func (*StudentRepository) GetStudentById(studentId string) (*models.Student, error) {
	student := &models.Student{}
	if err := global.DB.Where("student_id=?", studentId).First(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

// LinkStudentToClass 新增学生&班级关系表的记录
func (*StudentRepository) LinkStudentToClass(student *models.Student, class *models.Class) error {
	return global.DB.Model(student).Association("Classes").Append(class)
}

func (*StudentRepository) UnlinkStudentFromClass(student *models.Student, class *models.Class) error {
	return global.DB.Model(&student).Association("Classes").Delete(&class)
}
