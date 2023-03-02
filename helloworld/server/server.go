package main

import (
	"fmt"
	message "github.com/grpc-demo/helloworld/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:9999")
	// 创建一个grpc服务
	grpcServer := grpc.NewServer()
	// 自定义服务
	myService := message.MyService{}
	// 将自定义服务注册到grpc服务中
	message.RegisterMessageSenderServer(grpcServer, &myService)

	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("grpcServer.Serve err:", err)
		return
	}
}
