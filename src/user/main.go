package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	user "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// get etcd address from env
	etcdAddr, exists := os.LookupEnv("ETCD_ENDPOINT")

	if !exists {
		etcdAddr = "http://etcd:2379"
	}

	log.Println("ETCD_ENDPOINT is set to ", etcdAddr)

	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),
		retry.WithObserveDelay(20*time.Second),
		retry.WithRetryDelay(5*time.Second),
	)
	etcd, err := etcd.NewEtcdRegistryWithRetry([]string{etcdAddr}, retryConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", ":10001")

	svr := user.NewServer(
		new(UserServiceImpl),
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
