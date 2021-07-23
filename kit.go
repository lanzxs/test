package main

import (
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"net/http"
	EndPoint1 "test/EndPoint"
	"test/Service"
	"test/Transport"
)

// 服务发布

func main() {
	// 1.先创建我们最开始定义的Server/server.go
	s := Service.Server{}

	// 2.在用EndPoint/endpoint.go 创建业务服务
	hello := EndPoint1.MakeServerEndPointHello(s)
	Bye := EndPoint1.MakeServerEndPointBye(s)

	// 3.使用 kit 创建 handler
	// 固定格式
	// 传入 业务服务 以及 定义的 加密解密方法
	helloServer := httpTransport.NewServer(hello, Transport.HelloDecodeRequest, Transport.HelloEncodeResponse)
	sayServer := httpTransport.NewServer(Bye, Transport.ByeDecodeRequest, Transport.ByeEncodeResponse)

	// 使用http包启动服务
	go http.ListenAndServe("0.0.0.0:8000", helloServer)

	go http.ListenAndServe("0.0.0.0:8001", sayServer)
	fmt.Println("启动完成")
	select {}
}
