package repositories

import (
	"attendance_uniapp/global"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	DB *gorm.DB
}

func NewAttendanceRepository() *AttendanceRepository {
	return &AttendanceRepository{DB: global.DB}
}
