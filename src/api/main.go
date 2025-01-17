package main

import (
	"context"
	"log"
	"os"

	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user"
	"github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var etcdAddr string

func main() {

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



	h := server.New(
		server.WithHostPorts(":8888"),
	)


	userclient := userservice.MustNewClient(
		"user",
		client.WithResolver(r),
	)

	h.GET("/hello", func(ctx context.Context, req *app.RequestContext) {
		req.String(200, "Hello, World!")
	})

	api := h.Group("/api")
	user := api.Group("/user")
	user.GET("/login", func(c context.Context, ctx *app.RequestContext) {
		loginHandler(c, ctx, userclient)
	})


	h.Spin()
}

func loginHandler(ctx context.Context, req *app.RequestContext, client userservice.Client) {
	rpcreq := &user.LoginReq{Email: "rpc_call@test.com", Password: "rpc_call_passwprd"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.Login(ctx, rpcreq)
	cancel()
	if err != nil {
		log.Println("On error:", err)
		req.JSON(500, "Internal Server Error")
		return
	}
	log.Println(resp)
	req.JSON(200, resp)
}
