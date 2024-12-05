package utils

import (
	"fmt"
	"testing"
)

// 用于测试utils包的函数
// 使用 go test -v .\utils\ 命令运行测试

func TestGenerateJWT(t *testing.T) {
	JWT, _ := GenerateJWT("202200202055", "student")
	fmt.Println(JWT)
}
