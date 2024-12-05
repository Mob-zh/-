package services

import (
	"attendance_uniapp/models"
	"attendance_uniapp/repositories"
)

type StudentService struct {
	StudentRepo *repositories.StudentRepository
	ClassRepo   *repositories.ClassRepository
}

func NewStudentService(studentRepo *repositories.StudentRepository, classRepo *repositories.ClassRepository) *StudentService {
	return &StudentService{
		StudentRepo: studentRepo,
		ClassRepo:   classRepo,
	}
}

func (StudentServ *StudentService) GetStudentByIdService(studentId string) (*models.Student, error) {
	queryForStudent, err := StudentServ.StudentRepo.GetStudentById(studentId)
	if err != nil {
		return nil, err
	}
	return queryForStudent, nil
}

func (StudentServ *StudentService) EnrollStudentInClassService(studentId string, classId string) error {
	student := &models.Student{StudentId: studentId}
	class := &models.Class{ClassId: classId}
	// 由于studentId直接从JWT中获取，所以不需要检查studentId是否存在
	// _, err := StudentServ.StudentRepo.GetStudentById(studentId)

	// 检查classId是否存在
	if _, err := StudentServ.ClassRepo.GetClassById(classId); err != nil {
		return err
	}
	if err := StudentServ.StudentRepo.LinkStudentToClass(student, class); err != nil {
		return err
	}
	return nil
}

func (StudentServ *StudentService) StudentQuitFromClassService(studentId string, classId string) error {
	return StudentServ.StudentRepo.UnlinkStudentFromClass(&models.Student{StudentId: studentId}, &models.Class{ClassId: classId})
}
