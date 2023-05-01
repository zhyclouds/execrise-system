package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

// GetMd5
// 获取字符串的 md5 值
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("execrise-system")

// GenerateToken
// 生成 token
func GenerateToken(identity, name string) (string, error) {
	userClaim := UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("AnalyseToken Error: %v", claims.Valid)
	}
	return userClaim, nil
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Zhang HaoYu <310123665@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码"
	e.HTML = []byte("您的验证码是<b>" + code + "</b>")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "310123665@qq.com", "vplxaeoiawtpbjda", "smtp.qq.com"))
	if err != nil {
		return err
	}
	return nil
}

// GetUUid
// 获取 uuid
func GetUUid() string {
	return uuid.NewV4().String()
}

// GetRand
// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
