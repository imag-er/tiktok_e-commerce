package main

import (
	"context"
	"src/api/def"
	"src/api/handler"
	"src/api/middleware"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	def.InitAll()

	h := server.New(
		server.WithHostPorts(":8888"),
	)

	// cors中间件
	h.Use(middleware.InitCORSMiddleware())

	// jwt middleware
	authMiddleware := middleware.InitJWTMiddleware()

	// 路由注册
	h.Any("/hello", func(ctx context.Context, req *app.RequestContext) {
		req.JSON(consts.StatusOK, utils.H{"message": "hello world"})
	})

	/// 登录接口
	h.POST("/login", authMiddleware.LoginHandler)
	h.POST("/register", handler.Register);
	/// api jwt
	auth := h.Group("auth")
	/// 刷新jwt
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	h.Spin()
}
