package main

import (
	"context"
	"github.com/micro/micro/v3/service"
	pb "github.com/socylx/laracom/demo-service/proto/demo"
	"log"
	"time"
)

func main() {
	srv := service.New()
	srv.Init()

	client := pb.NewDemoService("demo", srv.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}
