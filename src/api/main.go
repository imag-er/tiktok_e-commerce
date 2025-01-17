package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	api "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user"
	clients "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
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


	h := server.New(
		server.WithHostPorts(":8888"),
	)

	h.GET("/hello", func(ctx context.Context, req *app.RequestContext) {
		req.String(200, "Hello, World!")
	})


	api := h.Group("/api")
	user := api.Group("/user")
	user.GET("/login", loginHandler)
	user.GET("/hello", func(ctx context.Context, req *app.RequestContext) {
		req.JSON(200, "Hello, World!")
	})

	h.Spin()
}

func loginHandler(ctx context.Context, req *app.RequestContext) {
	etcd, err := etcd.NewEtcdResolver([]string{etcdAddr})

	// 创建user服务的客户端
	client, err := clients.NewClient(
		"user",
		client.WithResolver(etcd),
	)

	if err != nil {
		log.Fatal(err)
	}
	
	rpcreq := &api.LoginReq{Email: "rpc_call@test.com", Password: "rpc_call_passwprd"}
	resp, err := client.Login(context.Background(), rpcreq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	req.JSON(200, resp)
}
