package jwt

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

func generateJwt(claims jwtgo.MapClaims, expiration time.Duration, secret string) (string, error) {
	uuid := utils.UUID()
	claims["uuid"] = uuid
	claims["iat"] = time.Now().Unix()
	claims["exp"] = expiration

	return generateToken(claims, secret)
}

func generateToken(claims jwtgo.Claims, secretKey string) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(token string, secretKey string) (*jwtgo.Token, jwtgo.MapClaims, error) {
	parsedToken, err := jwtgo.Parse(token, func(t *jwtgo.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}
	return parsedToken, parsedToken.Claims.(jwtgo.MapClaims), nil
}
