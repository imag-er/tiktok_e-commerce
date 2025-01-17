package main

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	api "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user"
	clients "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	h := server.New(
		server.WithHostPorts(":8888"),
	)

	api := h.Group("/api")
	user := api.Group("/user")
	user.GET("/login", loginHandler)

	h.Spin()
}

func loginHandler(ctx context.Context, req *app.RequestContext) {
	etcd, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})

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

}
