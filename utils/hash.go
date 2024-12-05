package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(plainPwd string, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
func GenerateClassId(className string, classTime string, teacherId string) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().UnixNano()
	// 将输入字符串和时间戳拼接
	data := fmt.Sprintf("%s-%d", className+classTime+teacherId, timestamp)
	// 使用 SHA256 生成哈希
	hash := sha256.Sum256([]byte(data))
	// 返回前 5 位
	return hex.EncodeToString(hash[:])[:5], nil
}
