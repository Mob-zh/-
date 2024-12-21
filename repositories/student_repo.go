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
	if err := global.DB.Where("student_id=?", studentId).First(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (StudentReno *StudentRepository) ChangeStudentPwdById(studentId string, newPwd string) error {
	return StudentReno.DB.Model(&models.Student{}).Where("student_id=?", studentId).Update("student_pwd", newPwd).Error
}
