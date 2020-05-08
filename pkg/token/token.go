package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

type Context struct {
	Keys string
}

func secreFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token
	token, err := jwt.Parse(tokenString, secreFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Keys = claims["keys"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string

	// Parse the header to get the token part
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin  config if the secret isn't specified
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	now := time.Now()
	// 120分钟之后
	thirty, _ := time.ParseDuration("87600h")
	thirtyMinutes := now.Add(thirty)

	//fmt.Println("开始时间",now.Unix())
	//fmt.Println("结束时间",thirtyMinutes.Unix())

	// The token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"keys": c.Keys,
		"iss":  "52bug.me",
		"nbf":  now.Unix(),
		"iat":  now.Unix(),
		"exp":  thirtyMinutes.Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))
	return

}
