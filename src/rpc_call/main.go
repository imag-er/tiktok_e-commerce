package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user"
	"github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"os"
	"time"
)

func main() {
	etcdAddr := os.Getenv("ETCD_ENDPOINT")
	log.Println("ETCD_ENDPOINT: ", etcdAddr)

	etcd, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	for {

		// 创建user服务的客户端
		client := userservice.MustNewClient(
			"user",
			client.WithResolver(etcd),
		)
	
		rpcreq := &user.LoginReq{Email: "rpc_call@test.com", Password: "rpc_call_passwprd"}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		resp, err := client.Login(ctx, rpcreq)
		cancel()
	
		if err != nil {
			log.Println("On error:", err)
			continue
		}
		log.Println(resp)

		time.Sleep(5 * time.Second)
	}

}
