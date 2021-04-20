package server

import (
	"context"
	"fmt"
	"github.com/maomaocode/codebase/grpc/proto"
)

type HelloServer struct {
}

func (h *HelloServer) HelloList(req *proto.HelloReq, server proto.Hello_HelloListServer) error {
	for i := 0; i < 5; i++ {
		_ = server.Send(&proto.HelloRes{Msg: fmt.Sprintf("server got %d %s", i, req.Name)})
	}
	return nil
}

func (h *HelloServer) Hello(context context.Context, req *proto.HelloReq) (*proto.HelloRes, error) {
	return &proto.HelloRes{Msg: fmt.Sprintf("server got %s", req.Name)}, nil
}
