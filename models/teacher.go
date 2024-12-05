package models

import "attendance_uniapp/initializer"

/*
教师表模型
*/

type Teacher struct {
	TeacherId   string `gorm:"primaryKey;type:varchar(12);not null" json:"teacher_id"` // 教师工号，主键且唯一
	TeacherName string `gorm:"type:varchar(50);not null" json:"teacher_name"`          // 教师姓名，非空
	TeacherPwd  string `gorm:"type:varchar(60);not null" json:"teacher_pwd"`           // 教师密码，非空
}

func GetTeacherById(teacherId string) (*Teacher, error) {
	teacher := &Teacher{}
	if err := initializer.DB.Where("teacher_id = ?", teacherId).First(teacher).Error; err != nil {
		return nil, err
	}
	return teacher, nil
}
