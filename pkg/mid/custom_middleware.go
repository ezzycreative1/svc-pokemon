package mid

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

func ApiKeyMiddleware(apiKey string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper: func(echo.Context) bool {
			return apiKey == ""
		},
		KeyLookup:  "header:X-API-Key",
		AuthScheme: "",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == apiKey, nil
		},
		ErrorHandler: func(err error, c echo.Context) error {
			response := map[string]interface{}{
				"message": "unauthorized",
				"data":    nil,
				"error":   err.Error(),
			}
			return c.JSON(http.StatusUnauthorized, response)
		},
	})
}

type UserAuth struct {
	Id int64
}

// GenerateToken ..
func GenerateToken(email string, userid int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userid
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// TokenValid ..
func TokenValid(c echo.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

// ExtractToken ..
func ExtractToken(c echo.Context) string {
	keys := c.Request().URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request().Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID ..
func ExtractTokenID(c echo.Context) (int64, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return int64(uid), nil
	}
	return 0, nil
}

//Pretty display the claims
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
