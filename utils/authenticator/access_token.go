package authenticator

import (
	"fmt"
	"time"

	"api/config"
	"api/model"

	"github.com/golang-jwt/jwt/v4"
)

type AccessToken interface {
	CreateAccessToken(cred *model.UserCredential) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type accessToken struct {
	cfg config.TokenConfig
}

func (t *accessToken) CreateAccessToken(cred *model.UserCredential) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.cfg.AccessTokenLifeTime)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.cfg.ApplicationName,
		},
		Username: cred.Username,
		Email:    cred.Email,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		t.cfg.JwtSigningMethod,
		claims,
	)

	return token.SignedString([]byte(t.cfg.JwtSignatureKey))
}

func (t *accessToken) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD")
		} else if method != t.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("INVALID SIGNING METHOD %v ", method)
		}
		return t.cfg.JwtSignatureKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		cfg: config,
	}
}
