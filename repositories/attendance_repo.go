package repositories

import (
	"attendance_uniapp/initializer"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	DB *gorm.DB
}

func NewAttendanceRepository() *AttendanceRepository {
	return &AttendanceRepository{DB: initializer.DB}
}
