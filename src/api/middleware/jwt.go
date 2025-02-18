package middleware

import (
	"context"
	"log"
	"src/api/def"
	"src/kitex_gen/user"
	"src/kitex_gen/user/userservice"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"

	"github.com/hertz-contrib/jwt"
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
			def.IdentityKey: v.UserId,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(ctx context.Context, c *app.RequestContext) interface{} {
	claims := jwt.ExtractClaims(ctx, c)
	return &def.User{
		UserId: claims[def.IdentityKey].(uint32),
	}

}

func authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var loginVals def.Login
	if err := c.BindAndValidate(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	req := user.LoginReq{
		Email:    loginVals.Email,
		Password: loginVals.Password,
	}

	resp, err := userservice.MustNewClient("user", client.WithResolver(def.EtcdResolver)).Login(ctx, &req)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return &def.User{
		UserId: resp.UserId,
	}, nil
}

func authorizator(data interface{}, ctx context.Context, c *app.RequestContext) bool {
	// TODO: replace this with your own authorization logic
	// TODO: casbin

	// default impl:
	if v, ok := data.(*def.User); ok && v.UserId <= 100000 {
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
