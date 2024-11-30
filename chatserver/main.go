package main

import (
	pb "blockchain-api/proto/chatserver"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

type chatServer struct {
	pb.UnimplementedChatServerServer
	mu     sync.Mutex
	client map[string]pb.ChatServer_MessageServer
}

func NewChatServer() *chatServer {
	return &chatServer{
		client: make(map[string]pb.ChatServer_MessageServer),
	}
}

func (cs *chatServer) Message(stream pb.ChatServer_MessageServer) error {
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("message received")

		switch message.GetMessageType() {

		case pb.MessageType_User:
			fmt.Println("registering new user")
			cs.mu.Lock()
			cs.client[message.Username] = stream
			cs.mu.Unlock()

		case pb.MessageType_Hello:
			fmt.Println("Message recived from", message.Username)
			for _, clientStream := range cs.client {
				clientStream.Send(&pb.ChatMessage{
					Username:    message.Username,
					MessageType: pb.MessageType_Hello,
				})
			}

		case pb.MessageType_Quit:
			fmt.Println("Removing user", message.Username)
			cs.mu.Lock()
			delete(cs.client, message.Username)
			cs.mu.Unlock()

		default:
			fmt.Println("default client")
		}
		fmt.Println(message.MessageType)

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterChatServerServer(s, NewChatServer())

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
