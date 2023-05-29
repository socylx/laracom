package main

import (
	"context"
	"github.com/micro/micro/v3/service"
	pb "github.com/socylx/laracom/demo-service/proto/demo"
	"log"
)

type DemoServiceHandler struct {
}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func main() {
	srv := service.New(service.Name("demo"))
	srv.Init()

	_ = pb.RegisterDemoServiceHandler(srv.Server(), &DemoServiceHandler{})
	if err := srv.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
