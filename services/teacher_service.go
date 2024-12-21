package services

import (
	"attendance_uniapp/models"
	"attendance_uniapp/repositories"
)

type TeacherService struct {
	TeacherRepo *repositories.TeacherRepository
}

func NewTeacherService(TeacherRepo *repositories.TeacherRepository) *TeacherService {
	return &TeacherService{TeacherRepo: TeacherRepo}
}

func (TeacherServ *TeacherService) GetTeacherByIdService(teacherId string) (*models.Teacher, error) {
	return TeacherServ.TeacherRepo.GetTeacherById(teacherId)
}

func (TeacherServ *TeacherService) ChangeTeacherPwdByIdService(teacherId string, newPwd string) error {
	return TeacherServ.TeacherRepo.ChangeTeacherPwdById(teacherId, newPwd)
}
