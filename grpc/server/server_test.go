package server

import (
	"context"
	"github.com/maomaocode/codebase/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"testing"
)

var (
	port = "9000"
)

func RunGrpcServer() error {
	server := grpc.NewServer()
	proto.RegisterHelloServer(server, &HelloServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	return server.Serve(lis)
}

func GetClient() proto.HelloClient {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	return proto.NewHelloClient(conn)
}

func SayHello() error {
	client := GetClient()
	req := &proto.HelloReq{Name: "maomao"}

	res, err := client.Hello(context.Background(), req)
	if err != nil {
		return err
	}
	log.Printf("client.Hello resp: %s", res.Msg)
	return nil
}

func SayHelloList() error {
	client := GetClient()
	req := &proto.HelloReq{Name: "maomao"}
	stream, err := client.HelloList(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("client.Hello resp: %s", res.Msg)
	}
	return nil
}

func TestHelloServer_Hello(t *testing.T) {
	go RunGrpcServer()
	if err := SayHello(); err != nil {
		return
	}

	if err := SayHelloList(); err != nil {
		return
	}
}
