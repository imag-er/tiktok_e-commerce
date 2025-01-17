package main

import (
	server "github.com/cloudwego/kitex/server"
	user "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"

	"log"
)

func main() {
	etcd, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(etcd),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
