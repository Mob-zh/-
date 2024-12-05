package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	StudentID  string
	StudentPwd string
}

func main() {
	// 创建一个 Student 对象，注意传递的是指针
	student := &Student{
		StudentID:  "S12345",
		StudentPwd: "password123",
	}

	// 获取 student 的反射对象，传递指针
	studentValue := reflect.ValueOf(student).Elem()

	// 打印类型和值
	fmt.Println("Type:", studentValue.Type()) // Output: Type: main.Student
	fmt.Println("Value:", studentValue)       // Output: Value: {S12345 password123}

	// 获取字段并修改
	studentPwdField := studentValue.FieldByName("StudentPwd")

	// 修改密码字段
	studentPwdField.SetString("newPassword456")

	student.StudentPwd = "ada"
	// 打印修改后的值
	fmt.Println("Modified Value:", studentValue) // Output: {S12345 newPassword456}
}
