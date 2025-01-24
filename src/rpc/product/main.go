package main

import (
	"log"
	product "src/kitex_gen/product/productcatalogservice"
)

func main() {
	svr := product.NewServer(new(ProductCatalogServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
