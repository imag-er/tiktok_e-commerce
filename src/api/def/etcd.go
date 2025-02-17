package def

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"os"
)

var EtcdResolver discovery.Resolver

func initEtcdService() {
	var etcdAddr string

	// get etcd address from env
	etcdAddr, exists := os.LookupEnv("ETCD_ENDPOINT")

	if !exists {
		etcdAddr = "http://127.0.0.1:2379"
	}

	log.Println("ETCD_ENDPOINT is set to ", etcdAddr)

	var err error
	EtcdResolver, err = etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}

}
