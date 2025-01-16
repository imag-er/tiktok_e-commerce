package main

import (
    "context"
	api "../user/kitex_gen/user"
    clients "../user/kitex_gen/user/userservice"
    "log"
)

func main() {
	// 创建user服务的客户端
    client, err := clients.NewClient("user")
    if err != nil {
        log.Fatal(err)
    }
    // req := &cli.EchoRequest{Msg: "Hello, Kitex!"}
	req := &api.LoginReq{Email: "rpc_call@test.com", Password: "rpc_call_passwprd"}
    resp, err := client.Login(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(resp)
}
