package repositories

import (
	"attendance_uniapp/global"
	"attendance_uniapp/models"
	"gorm.io/gorm"
)

type StudentClassRepository struct {
	DB *gorm.DB
}

func NewStudentClassRepository() *StudentClassRepository {
	return &StudentClassRepository{DB: global.DB}
}

// LinkStudentToClass 新增学生&班级关系表的记录
func (*StudentClassRepository) LinkStudentToClass(student *models.Student, class *models.Class) error {
	return global.DB.Model(student).Association("Classes").Append(class)
}

func (*StudentClassRepository) UnlinkStudentFromClass(student *models.Student, class *models.Class) error {
	return global.DB.Model(student).Association("Classes").Delete(class)
}
