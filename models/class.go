package models

import (
	"time"
)

/*
	班级表模型
*/

type Class struct {
	ClassId            string             `gorm:"primaryKey;type:varchar(8);not null;" json:"class_id"`                      // 班级ID，非空，唯一主键
	ClassName          string             `gorm:"type:varchar(50);not null" json:"class_name"`                               // 班级名称，非空
	TeacherId          string             `gorm:"type:varchar(12);not null;index" json:"teacher_id"`                         // 教师ID，非空，外键
	CourseName         string             `gorm:"type:varchar(50);not null" json:"course_name"`                              // 课程名，非空
	CourseId           string             `gorm:"type:varchar(20);not null;index" json:"course_id"`                          // 课程ID，非空，外键
	ClassTime          string             `gorm:"type:varchar(20);not null" json:"class_time"`                               // 上课时间，非空
	TeacherName        string             `gorm:"type:varchar(100);not null" json:"teacher_name"`                            // 教师姓名，非空
	IsClassChecking    bool               `gorm:"not null;type:boolean" json:"is_class_checking"`                            // 班级考勤状态，非空
	CheckingEndTime    time.Time          `gorm:"type:datetime;not null" json:"checking_endtime"`                            // 考勤结束时间，非空
	TotalCheckingTimes int                `gorm:"not null" json:"total_checking_times"`                                      // 总签到次数，非空
	Classroom          string             `gorm:"type:varchar(30);not null" json:"classroom"`                                // 教室，非空
	Students           []Student          `gorm:"many2many:student_classes;" json:"students"`                                // 自动创建中间表 student_classes
	AttendanceRecords  []AttendanceRecord `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE;" json:"attendance_records"` // 外键与级联删除
}
