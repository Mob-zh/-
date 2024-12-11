package main

/*
 * 迁移数据库
go run migrations/migrate.go
*/

import (
	"attendance_uniapp/global"
	"attendance_uniapp/initializer"
	"attendance_uniapp/models"
	"log"
)

var modelList = []interface{}{
	&models.Student{},          // Student 模型
	&models.Teacher{},          // Teacher 模型
	&models.Class{},            // Class 模型
	&models.Course{},           // Course 模型
	&models.AttendanceRecord{}, // AttendanceRecord 模型
}

func main() {
	initializer.Init()
	// 自动迁移所有模型
	for _, model := range modelList {
		if err := global.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("Migration complete!")
}
