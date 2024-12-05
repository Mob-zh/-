package services

import "attendance_uniapp/repositories"

type AttendanceService struct {
	AttendanceRepo *repositories.AttendanceRepository
}

func NewAttendanceService(AttendanceRepo *repositories.AttendanceRepository) *AttendanceService {
	return &AttendanceService{AttendanceRepo: AttendanceRepo}
}
