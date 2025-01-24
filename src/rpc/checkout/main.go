package main

import (
	"log"
	checkout "src/kitex_gen/checkout/checkoutservice"
)

func main() {
	svr := checkout.NewServer(new(CheckoutServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
