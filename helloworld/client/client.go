package main

import (
	"context"
	"fmt"
	message "github.com/grpc-demo/helloworld/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := message.NewMessageSenderClient(conn)
	// 调用服务端方法
	resp, err := client.Send(context.Background(), &message.MessageRequest{SaySomething: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	s := resp.String()
	fmt.Println(s)
}
