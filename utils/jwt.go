package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userId string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 7 * 24).Unix(),
	})
	// 签名密钥为"SZTU"
	signedToken, err := token.SignedString([]byte("SZTU"))
	return "Bearer " + signedToken, err
}

func ParseJWT(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("SZTU"), nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), claims["role"].(string), nil
	}
	return "", "", errors.New("invalid token")

}
