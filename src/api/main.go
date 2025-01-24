package main

import (
	"context"
	"log"
	"os"

	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/hertz-contrib/cors"

	"github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user"
	"github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	h := server.New(
		server.WithHostPorts(":8888"),
	)
	// 中间件注册
	h.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type"},
		AllowMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// rpc服务发现
	r := initEtcdService()
	userclient := userservice.MustNewClient("user", client.WithResolver(r))

	// 路由注册
	h.Any("/hello", func(ctx context.Context, req *app.RequestContext) {
		req.JSON(consts.StatusOK, utils.H{"message": "hello world"})
	})

	api := h.Group("/api")
	user := api.Group("/user")
	user.POST("/login", func(c context.Context, ctx *app.RequestContext) {
		loginHandler(c, ctx, userclient)
	})

	h.Spin()
}

func loginHandler(_ context.Context, httpctx *app.RequestContext, client userservice.Client) {

	rpcreq := &user.LoginReq{Email: httpctx.PostForm("email"), Password: httpctx.PostForm("password")}

	rpcctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.Login(rpcctx, rpcreq)
	cancel()

	if err != nil {
		log.Println("On error:", err)
		httpctx.JSON(500, "Internal Server Error")
		return
	}
	log.Println(resp)
	httpctx.JSON(200, resp)
}

func initEtcdService() discovery.Resolver {
	var etcdAddr string

	// get etcd address from env
	etcdAddr, exists := os.LookupEnv("ETCD_ENDPOINT")

	if !exists {
		etcdAddr = "http://etcd:2379"
	}

	log.Println("ETCD_ENDPOINT is set to ", etcdAddr)

	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	return r
}
