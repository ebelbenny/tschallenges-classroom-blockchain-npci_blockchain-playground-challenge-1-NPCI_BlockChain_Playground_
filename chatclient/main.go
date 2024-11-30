package main

import (
	pb "blockchain-api/proto/chatserver"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	userName := os.Getenv("USERNAME")
	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	fmt.Println("dialing")
	time.Sleep(10 * time.Second)
	chatServerClient := pb.NewChatServerClient(conn)
	ctx := context.Background()
	stream, err := chatServerClient.Message(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = stream.Send(&pb.ChatMessage{
		Username:    userName,
		MessageType: pb.MessageType_User,
	})

	go ReceiveMsg(stream, userName)
	fmt.Println("sending")
	if err != nil {
		fmt.Println(err)
	}

	err = stream.Send(&pb.ChatMessage{
		Username:    userName,
		MessageType: pb.MessageType_Hello,
	})

	fmt.Println("sending")
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(50 * time.Second)
}

func ReceiveMsg(stream pb.ChatServer_MessageClient, username string) {

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return
		}
		fmt.Println("message received")

		switch message.GetMessageType() {
		case pb.MessageType_Hello:
			fmt.Println(message.Username)

		default:
			fmt.Println("default client")
		}


		fmt.Println(message.MessageType)
	}

}
