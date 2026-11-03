package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"pledge-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string) (string, error) {
	if len(config.Config.Jwt.SecretKey) < 32 {
		return "", errors.New("JWT secret is not configured")
	}
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username, "iat": now.Unix(), "exp": now.Add(time.Duration(config.Config.Jwt.ExpireTime) * time.Second).Unix(), "jti": hex.EncodeToString(nonce),
	})
	return token.SignedString([]byte(config.Config.Jwt.SecretKey))
}

func ParseToken(raw, secret string) (string, error) {
	if len(secret) < 32 {
		return "", errors.New("JWT secret is not configured")
	}
	token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected JWT signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if token == nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}
	name, ok := claims["username"].(string)
	if !ok || name == "" {
		return "", errors.New("missing username")
	}
	if _, ok := claims["exp"]; !ok {
		return "", errors.New("missing expiry")
	}
	return name, nil
}
