package mid

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ezzycreative1/svc-pokemon/pkg/convert"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

// Custom JWT middleware
func JwtMiddleware(mayangSecretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Get token from header
			authorizationHeader := ctx.Request().Header.Get("Authorization")

			// Check if string has token
			if !strings.Contains(authorizationHeader, "Bearer") {
				response := map[string]interface{}{
					"message": "Bad request",
					"data":    nil,
					"error":   "invalid token or missing token",
				}
				return ctx.JSON(http.StatusBadRequest, response)
			}

			// Split token string
			tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

			// Parse token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("signing method invalid")
				} else if method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("signing method invalid")
				}

				return []byte(mayangSecretKey), nil
			})

			// Check if there is error
			if err != nil {
				response := map[string]interface{}{
					"message": "Bad request",
					"data":    nil,
					"error":   err.Error(),
				}
				return ctx.JSON(http.StatusBadRequest, response)
			}

			// Save user_id and branch_id to context
			claims := token.Claims.(jwt.MapClaims)
			user_id := fmt.Sprintf("%v", claims["id"])

			// Set data to struct
			UserAuthData := UserAuth{
				Id: convert.StrToInt64(user_id, 0),
			}

			// Set to Context
			ctx.Set("UserAuth", &UserAuthData)

			// Skip to router
			return next(ctx)
		}

	}
}
