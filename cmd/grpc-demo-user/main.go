package main

import (
	"sync"

	"github.com/wuyadong1990/grpc-demo-user/cinit"
	//"github.com/wuyadong1990/grpc-demo-user/internal/impl"
	"github.com/wuyadong1990/grpc-demo-user/internal/server/grpc"
	"github.com/wuyadong1990/grpc-demo-user/internal/server/http"
)

const (
	SN = "srv-user" // 定义services名称
)

func main() {

	// 初始化,选着需要的组件
	cinit.InitOption(SN, cinit.Trace, cinit.MySQL, cinit.Redis)
	//cinit.GormDB.AutoMigrate(&(impl.User))
	var wg sync.WaitGroup

	wg.Add(1)
	go grpc.Serve(&wg, "5000")

	wg.Add(1)
	go http.Serve(&wg, "5000", "8080")

	wg.Wait()
}
