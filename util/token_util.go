package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tejashwinn/splitwise/types"
)

type JwtUtil struct {
	secretKey              []byte
	issuer                 string
	jwtAccessTokenExpMin   time.Duration
	jwtRefreshTokenExpHour time.Duration
}

func NewJwtUtil(cfg *types.Config) *JwtUtil {
	return &JwtUtil{
		secretKey:              cfg.Jwt.SecretKey,
		issuer:                 cfg.Jwt.Issuer,
		jwtAccessTokenExpMin:   time.Duration(rand.Int31n(cfg.Jwt.JwtAccessTokenExpMin)),
		jwtRefreshTokenExpHour: time.Duration(rand.Int31n(cfg.Jwt.JwtRefreshTokenExpHour)),
	}
}

func (jwtUtil *JwtUtil) GenerateToken(user *types.User) (string, string, error) {
	// Set token expiration times
	accessTokenExpiry := time.Now().Add(jwtUtil.jwtAccessTokenExpMin * time.Minute)
	refreshTokenExpiry := time.Now().Add(jwtUtil.jwtRefreshTokenExpHour * time.Hour)

	// Create access token
	accessClaims := jwt.MapClaims{
		"sub": user.Id,
		"exp": accessTokenExpiry.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtUtil.secretKey)

	if err != nil {
		return "", "", err
	}

	// Create refresh token
	refreshClaims := jwt.MapClaims{
		"sub": user.Id,
		"exp": refreshTokenExpiry.Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtUtil.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (jwtUtil *JwtUtil) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtUtil.secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("invalid user ID")
	}

	return userID, nil
}

func (jwtUtil *JwtUtil) RefreshToken(refreshTokenString string) (string, error) {
	userID, err := jwtUtil.VerifyToken(refreshTokenString)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token")
	}

	newAccessTokenExpiry := time.Now().Add(15 * time.Minute)
	accessClaims := jwt.MapClaims{
		"sub": userID,
		"exp": newAccessTokenExpiry.Unix(),
	}
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	return newAccessToken.SignedString(jwtUtil.secretKey)
}
