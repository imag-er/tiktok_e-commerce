package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"time"
	"src/api/def"
)

func InitJWTMiddleware() (authMiddleware *jwt.HertzJWTMiddleware) {

	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: def.IdentityKey,

		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		PayloadFunc:     payloadFunc,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	if errInit := authMiddleware.MiddlewareInit(); errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*def.User); ok {
		return jwt.MapClaims{
			def.IdentityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(ctx context.Context, c *app.RequestContext) interface{} {
	claims := jwt.ExtractClaims(ctx, c)
	return &def.User{
		UserName: claims[def.IdentityKey].(string),
	}

}

func authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var loginVals def.Login
	if err := c.BindAndValidate(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	// TODO: replace this with your own authentication logic
	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return &def.User{
			UserName: userID,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func authorizator(data interface{}, ctx context.Context, c *app.RequestContext) bool {
	// TODO: replace this with your own authorization logic

	// TODO: casbin
	if v, ok := data.(*def.User); ok && v.UserName == "admin" {
		return true
	}

	return false
}

func unauthorized(ctx context.Context, c *app.RequestContext, code int, message string) {
	c.JSON(code, map[string]interface{}{
		"code":    code,
		"message": message,
	})
}
