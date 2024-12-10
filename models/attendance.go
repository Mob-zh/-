package models

type AttendanceRecord struct {
	StudentId   string `gorm:"primaryKey;type:varchar(12);not null" json:"student_id"`                                    // 联合主键，非空
	StudentName string `gorm:"type:varchar(50);not null" json:"student_name"`                                             // 非空字段
	ClassId     string `gorm:"primaryKey;type:varchar(5);not null;foreignKey:ClassId;references:ClassId" json:"class_id"` // 外键约束
	SigninTime  string `gorm:"primaryKey;type:datetime;not null" json:"signin_time"`                                      // 联合主键，非空
	IsSignin    bool   `gorm:"not null" json:"is_signin"`                                                                 // 非空字段
	SigninIndex int    `gorm:"not null" json:"signin_index"`                                                              // 签到索引（记录为第几次签到）
	Class       Class  `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE;" json:"class"`
}
