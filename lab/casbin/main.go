package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	// 加载模型和策略
	enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		fmt.Printf("NewEnforcer failed, error: %v\n", err)
		return
	}

	// 检查权限
	sub := "alice" // 请求的实体
	obj := "data1" // 请求的资源
	act := "read"  // 请求的操作

	if res, _ := enforcer.Enforce(sub, obj, act); res == true {
		fmt.Printf("%s can %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s cannot %s %s\n", sub, act, obj)
	}
}
