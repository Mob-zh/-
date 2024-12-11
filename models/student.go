package models

import (
	"attendance_uniapp/global"
)

/*
	学生表模型
*/

type Student struct {
	StudentId   string  `json:"student_id" gorm:"type:varchar(12);not null;comment:学生ID;primaryKey;"`
	StudentName string  `json:"student_name" gorm:"type:varchar(50);not null;comment:学生姓名;"`
	StudentPwd  string  `json:"student_pwd" gorm:"type:varchar(60);not null;comment:学生密码;"`
	Classes     []Class `gorm:"many2many:student_classes;" json:"classes"` // 自动创建中间表 student_classes
}

func GetStudentById(studentId string) (*Student, error) {
	student := &Student{}
	if err := global.DB.Where("student_id=?", studentId).First(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}
