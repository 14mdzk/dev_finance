package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	AccessTokenKey       string
	RefreshTokenKey      string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func NewTokenService(
	accessTokenKey string,
	refreshTokenKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *TokenService {
	return &TokenService{
		AccessTokenKey:       accessTokenKey,
		RefreshTokenKey:      refreshTokenKey,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
}

func (svc *TokenService) GenerateAccessToken(userID int) (string, time.Time, error) {
	exp := time.Now().Add(svc.AccessTokenDuration)
	key := []byte(svc.AccessTokenKey)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}

func (svc *TokenService) GenerateRefreshToken(userID int) (string, time.Time, error) {
	exp := time.Now().Add(svc.RefreshTokenDuration)
	key := []byte(svc.RefreshTokenKey)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}

func (svc *TokenService) VerifyAccessToken(tokenString string) (string, error) {
	sub, err := svc.verify(tokenString, svc.AccessTokenKey)
	return sub, err
}

func (svc *TokenService) VerifyRefreshToken(tokenString string) (string, error) {
	sub, err := svc.verify(tokenString, svc.RefreshTokenKey)
	return sub, err
}

func (svc *TokenService) verify(tokenString string, tokenKey string) (string, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sigining method: %v", token.Header["alg"])
		}

		return []byte(tokenKey), nil
	}

	token, err := jwt.Parse(tokenString, keyfunc)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sub := fmt.Sprint(claims["sub"])
		return sub, nil
	}

	return "", nil
}
