package services

import (
	"attendance_uniapp/models"
	"attendance_uniapp/repositories"
)

// ClassService 封装班级相关的业务逻辑
type ClassService struct {
	ClassRepo        *repositories.ClassRepository
	StudentRepo      *repositories.StudentRepository
	StudentClassReno *repositories.StudentClassRepository
}

// NewClassService 返回一个 ClassService 实例
func NewClassService(ClassRepo *repositories.ClassRepository, StudentReno *repositories.StudentRepository, StudentClassReno *repositories.StudentClassRepository) *ClassService {
	return &ClassService{ClassRepo: ClassRepo, StudentRepo: StudentReno, StudentClassReno: StudentClassReno}
}

func (ClassServ *ClassService) GetClassByIdService(classId string) (*models.Class, error) {
	return ClassServ.ClassRepo.GetClassById(classId)
}

func (ClassServ *ClassService) GetStudentListByClassIdService(classId string) ([]models.Student, error) {
	return ClassServ.ClassRepo.GetStudentInfoListByClassId(classId)
}

func (ClassServ *ClassService) GetClassListByStudentIdService(studentId string) ([]models.Class, error) {
	return ClassServ.ClassRepo.GetClassListByStudentId(studentId)
}

func (ClassServ *ClassService) DeleteClassByIdService(classId string) error {
	return ClassServ.ClassRepo.DeleteClassById(classId)
}

func (ClassServ *ClassService) StudentQuitFromClassService(studentId string, classId string) error {
	return ClassServ.StudentClassReno.UnlinkStudentFromClass(&models.Student{StudentId: studentId}, &models.Class{ClassId: classId})
}

func (ClassServ *ClassService) EnrollStudentInClassService(studentId string, classId string) error {
	student := &models.Student{StudentId: studentId}
	class := &models.Class{ClassId: classId}
	if err := ClassServ.StudentClassReno.LinkStudentToClass(student, class); err != nil {
		return err
	}
	return nil
}
