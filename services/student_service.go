package services

import (
	"attendance_uniapp/models"
	"attendance_uniapp/repositories"
)

type StudentService struct {
	StudentRepo    *repositories.StudentRepository
	ClassRepo      *repositories.ClassRepository
	AttendanceReno *repositories.AttendanceRepository
}

func NewStudentService(studentRepo *repositories.StudentRepository, classRepo *repositories.ClassRepository, attendanceReno *repositories.AttendanceRepository) *StudentService {
	return &StudentService{StudentRepo: studentRepo, ClassRepo: classRepo, AttendanceReno: attendanceReno}
}

func (StudentServ *StudentService) GetStudentByIdService(studentId string) (*models.Student, error) {
	queryForStudent, err := StudentServ.StudentRepo.GetStudentById(studentId)
	if err != nil {
		return nil, err
	}
	return queryForStudent, nil
}

func (StudentServ *StudentService) ChangeStudentPwdByIdService(studentId string, newPwd string) error {
	return StudentServ.StudentRepo.ChangeStudentPwdById(studentId, newPwd)
}
