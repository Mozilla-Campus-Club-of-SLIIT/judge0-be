package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: add rate limit
func GetBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("Invalid Authorization header format")
	}

	if parts[1] == "" {
		return "", errors.New("JWT token missing")
	}

	return parts[1], nil
}

// TODO: check expire and all
func ParseJWT(tokenStr string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse JWT claims")
	}

	return claims, nil
}

func GetRolesFromClaims(claims jwt.MapClaims) ([]string, error) {
	rawRoles, ok := claims["role"].([]interface{})
	if !ok {
		return nil, errors.New("roles missing in JWT claims")
	}

	var roles []string
	for _, r := range rawRoles {
		if roleStr, ok := r.(string); ok {
			roles = append(roles, roleStr)
		}
	}

	return roles, nil
}

func HasRequiredRole(userRoles, requiredRoles []string) bool {
	for _, u := range userRoles {
		for _, r := range requiredRoles {
			if u == r {
				return true
			}
		}
	}
	return false
}
