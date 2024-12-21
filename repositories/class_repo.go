package repositories

import "C"
import (
	"attendance_uniapp/global"
	"attendance_uniapp/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ClassRepository struct {
	DB     *gorm.DB
	Client *redis.Client
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

func (classReno *ClassRepository) GetStudentInfoListByClassId(classId string) ([]models.Student, error) {
	var class models.Class
	var studentInfoList []models.Student
	err := classReno.DB.Preload("Students").Order("student_id").First(&class, classId).Error
	//密码字段置为空,防止泄露
	for i := range class.Students {
		studentInfoList = append(studentInfoList, models.Student{StudentId: class.Students[i].StudentId, StudentName: class.Students[i].StudentName, StudentPwd: ""})
	}
	return studentInfoList, err
}

func (classReno *ClassRepository) GetClassListByStudentId(studentId string) ([]models.Class, error) {
	var student models.Student
	return student.Classes, classReno.DB.Preload("Classes").Order("class_id").First(&student, studentId).Error
}

func (classReno *ClassRepository) DeleteClassById(classId string) error {
	return global.DB.Where("class_id = ?", classId).Delete(&models.Class{}).Error
}

func (classReno *ClassRepository) GetStudentCountByClassId(classId string) int {
	count := classReno.DB.Model(&models.Student{}).
		Where("class_id = ?", classId).
		Association("Classes").Count()
	return int(count)
}
