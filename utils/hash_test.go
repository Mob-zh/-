package utils

import (
	"fmt"
	"testing"
)

func TestGenerateClassId(t *testing.T) {
	classId, _ := GenerateClassId("软工12班", "周一 5-6节", "BD3160812")
	fmt.Println("classId:", classId)
}
