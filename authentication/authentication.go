package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID string
	jwt.StandardClaims
}

func ParseToken(TokenString string, Jwtkey []byte) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	Token, err := jwt.ParseWithClaims(TokenString, Claims, func(Token *jwt.Token) (i interface{}, err error) {
		return Jwtkey, nil
	})
	return Token, Claims, err
}

//颁发token
func Setting(Info map[string]string, ExpireTime time.Time) (string, error) {
	ID := Info["ID"]
	JwtkeyString := Info["JwtkeyString"]

	var Jwtkey []byte = []byte(JwtkeyString)
	// ExpireTime := time.Now().Add(1 * 24 * time.Hour)
	Claims := &Claims{
		ID: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "sms",        // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	TokenString, err := Token.SignedString(Jwtkey)
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

//解析token
func Getting(Info map[string]string) (string, error) {
	TokenString := Info["TokenString"]
	JwtkeyString := Info["JwtkeyString"]

	var Jwtkey []byte = []byte(JwtkeyString)
	if TokenString == "" {
		return "", nil
	}
	Token, Claims, err := ParseToken(TokenString, Jwtkey)
	if err != nil || !Token.Valid {
		return "", err
	}
	return Claims.ID, nil
}
