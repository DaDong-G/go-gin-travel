package Jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//func main()  {
//	//tocken, _ := CreateToken([]byte("hhahsadfas"), "gjd", "haha", true)
//	claims, err := ParseToken("eyJhbGaciOiJIUzI1NiIsnR5cCI6IkpXVCJ9.eyJleHAsiojE1NzcwODQxNTUsImlzcyI6ImdqZCIsInVpZCI6ImhhaGEiLCJhZG1pbiI6dHJ1ZX0.qrM5lcLLeVuepin8CnPmnYBtkp1Yi0udjFJdvlbKM0o", []byte("hhahsadfas"))
//	//fmt.Println(tocken)
//	fmt.Println(claims)
//	fmt.Println(err)
//}

type jwtCustomClaims struct {
	jwt.StandardClaims
	// 追加自己需要的信息
	UserName   string `json:"uid"`
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(SecretKey []byte, issuer string, UserName string) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
