package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/joho/godotenv"
)

var API_SECRET string

type Claims struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
	jwt.RegisteredClaims
}

func init() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		PanicIfError(err)
		API_SECRET = MustGetEnv("API_SECRET")
	}
}
func GenerateToken(userId uint, roleId uint) (string, error) {
	tokenLifeSpan, err := strconv.Atoi(MustGetEnv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := &Claims{
		UserID: userId,
		RoleID: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(tokenLifeSpan))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(API_SECRET))
}

func GenerateResetPasswordToken(userId uint) (string, error) {
	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenClaims(c *gin.Context) (*Claims, error) {
	tokenString := ExtractToken(c)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
		}
		return []byte(API_SECRET), nil
	})

	if err != nil {
		return nil, exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
	}

	if !token.Valid {
		return nil, exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
	}

	return claims, nil
}

func UserRoleMustAdmin(c *gin.Context) {
	claims, err := ExtractTokenClaims(c)
	if err != nil {
		PanicIfError(err)
	}
	if claims.RoleID != uint(entity.IDRoleAdmin) {
		PanicIfError(exceptions.NewCustomError(http.StatusForbidden, "Only admin can manipulate data"))
	}
}
