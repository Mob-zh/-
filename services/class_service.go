package services

import (
	"attendance_uniapp/models"
	"attendance_uniapp/repositories"
)

// ClassService 封装班级相关的业务逻辑
type ClassService struct {
	ClassRepo   *repositories.ClassRepository
	StudentRepo *repositories.StudentRepository
}

// NewClassService 返回一个 ClassService 实例
func NewClassService(ClassRepo *repositories.ClassRepository) *ClassService {
	return &ClassService{ClassRepo: ClassRepo}
}

func (ClassServ *ClassService) GetClassByIdService(classId string) (*models.Class, error) {
	return ClassServ.ClassRepo.GetClassById(classId)
}

func (ClassServ *ClassService) GetStudentListByClassIdService(classId string) ([]*models.Student, error) {
	return ClassServ.ClassRepo.GetStudentListByClassId(classId)
}

func (ClassServ *ClassService) GetClassListByStudentIdService(studentId string) ([]*models.Class, error) {
	return ClassServ.ClassRepo.GetClassListByStudentId(studentId)
}

func (ClassServ *ClassService) DeleteClassByIdService(classId string) error {
	return ClassServ.ClassRepo.DeleteClassById(classId)
}
