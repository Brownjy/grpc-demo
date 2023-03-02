package message

import (
	"context"
	"log"
)

type MyService struct {
}

func (myService *MyService) Send(ctx context.Context, req *MessageRequest) (*MessageResponse, error) {
	log.Println("receive message:", req.GetSaySomething())
	resp := &MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}
func (myService *MyService) mustEmbedUnimplementedMessageSenderServer() {}
