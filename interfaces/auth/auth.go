package auth

import (
	"ca-mission/domain/model"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
)

// Token 作成関数
func CreateToken(user model.User) (string, error) {
	var err error

	// 鍵となる文字列(多分なんでもいい)
	secret := "secret"

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"name": user.Name,
		"iss":  "admin", // JWT の発行者が入る(文字列(__init__)は任意)
	})

	//Dumpを吐く
	spew.Dump(token)

	tokenString, err := token.SignedString([]byte(secret))

	// fmt.Println("-----------------------------")
	// fmt.Println("tokenString:", tokenString)

	return tokenString, err

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // CreateTokenにて指定した文字列を使う
	})
	if err != nil {
		return token, err
	}
	return token, nil
}

func ReadClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}
