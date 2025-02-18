package main

import (
	"log"
	"net"
	"os"
	
	user "src/kitex_gen/user/userservice"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
)

func main() {
	// get etcd address from env
	etcdAddr, exists := os.LookupEnv("ETCD_ENDPOINT")

	if !exists {
		etcdAddr = "http://127.0.0.1:2379"
		// etcdAddr = "http://etcd:2379"
	}

	log.Println("ETCD_ENDPOINT is set to ", etcdAddr)

	// connect to etcd
	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),
		retry.WithRetryDelay(5*time.Second),
	)
	etcd, err := etcd.NewEtcdRegistryWithRetry([]string{etcdAddr}, retryConfig)
	if err != nil {
		log.Fatalln("连接到etcd失败 ", err.Error())
	}

	// create server
	addr, _ := net.ResolveTCPAddr("tcp", ":10001")
	
	userservice := new(UserServiceImpl)
	userservice.db = InitGORM()

	svr := user.NewServer(
		userservice,
		server.WithServiceAddr(addr),
		server.WithRegistry(etcd),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "user",
			}),
	)

	err = svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
