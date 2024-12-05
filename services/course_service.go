package services

import (
	"attendance_uniapp/models"
	"attendance_uniapp/repositories"
)

type CourseService struct {
	CourseRepo *repositories.CourseRepository
}

func NewCourseService(CourseRepo *repositories.CourseRepository) *CourseService {
	return &CourseService{CourseRepo: CourseRepo}
}

func (CourseServ *CourseService) GetCourseByIdService(courseId string) (*models.Course, error) {
	return CourseServ.CourseRepo.GetCourseById(courseId)
}
