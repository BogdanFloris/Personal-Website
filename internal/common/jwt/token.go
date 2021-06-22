package jwt

import (
	"bogdanfloris-com/internal/common/config"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

type TokenType struct {
	string
	time.Duration
}

func AccessTokenType() TokenType {
	return TokenType{"ACCESS", 15 * time.Minute}
}

func RefreshTokenType() TokenType {
	return TokenType{"REFRESH", 24 * 7 * time.Hour}
}

type Token struct {
	token     string
	expiresAt time.Time
	tokenType TokenType
}

func NewToken(claims jwtgo.MapClaims, tokenType TokenType) (*Token, error) {
	exp := config.GetDuration(fmt.Sprintf("%s_TOKEN_DURATION", tokenType.string),
		time.Duration(time.Now().Add(tokenType.Duration).Unix()))
	secret := config.GetString(fmt.Sprintf("%s_TOKEN_SECRET", tokenType.string), "")

	t, err := generateJwt(claims, exp, secret)
	if err != nil {
		return nil, err
	}

	token := &Token{t, time.Now().Add(exp), tokenType}

	return token, nil
}

func (token *Token) VerifyToken() (jwtgo.MapClaims, error) {
	t, claims, err := verifyToken(token.token, config.GetString(fmt.Sprintf(
		"%s_TOKEN_SECRET", token.tokenType.string), ""))
	if err != nil {
		return nil, err
	}
	if _, ok := t.Claims.(jwtgo.Claims); !ok && !t.Valid {
		return nil, jwtgo.ErrInvalidKey
	}
	return claims, nil
}
