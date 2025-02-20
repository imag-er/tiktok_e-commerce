package main

import (
	"log"
	"os"
	"time"
	product "src/kitex_gen/product/productcatalogservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
	"net"
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
	addr, _ := net.ResolveTCPAddr("tcp", ":10002")
	
	productservice := new(ProductCatalogServiceImpl)
	productservice.db = InitGORM()

	svr := product.NewServer(
		productservice,
		server.WithServiceAddr(addr),
		server.WithRegistry(etcd),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "product",
			}),
	)

	err = svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}


