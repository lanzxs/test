package main

import (
	"github.com/micro-plat/hydra"
	//"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

func main() {

	//创建app
	/*appHttp := hydra.NewApp(
		hydra.WithServerTypes(http.API),
	)

	//注册服务
	appHttp.API("/hello", func(ctx hydra.IContext) interface{} {
		return "hello world"
	})

	appHttp.Start()*/

	//创建app
	app := hydra.NewApp(
		hydra.WithServerTypes(rpc.RPC),
	)

	//注册服务
	app.RPC("/hello", func(ctx hydra.IContext) interface{} {
		return "hello world"
	})

	//启动app
	app.Start()
}
