package middleware

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func InitJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("your_secret_key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals struct {
				Username string `form:"username" json:"username" binding:"required"`
				Password string `form:"password" json:"password" binding:"required"`
			}
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "password") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Muhammad",
					FirstName: "Hanafi",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:    "cookie:token_jwt",
		SendCookie:     true,
		SecureCookie:   false,
		CookieHTTPOnly: true,
		CookieName:     "token_jwt",
		CookieSameSite: http.SameSiteDefaultMode,
		TimeFunc:       time.Now,
	})
	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
