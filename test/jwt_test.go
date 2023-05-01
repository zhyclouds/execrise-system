package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("execrise-system")

// 生成 token
func TestGenerateToken(t *testing.T) {
	userClaim := UserClaims{
		Identity:       "admin",
		Name:           "admin",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tokenString)
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6ImFkbWluIiwibmFtZSI6ImFkbWluIn0.g6RXqKhyGY4aNcF31an2v_MScfkL39L5D7qMUPTKWk0
// 解析 token
func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6ImFkbWluIiwibmFtZSI6ImFkbWluIn0.g6RXqKhyGY4aNcF31an2v_MScfkL39L5D7qMUPTKWk0"
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaim)
	}
}
