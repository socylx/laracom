package main

import (
	"context"
	"github.com/micro/micro/v3/service"
	pb "github.com/socylx/laracom/user-service/proto/user"
	"log"
	"time"
)

func main() {
	srv := service.New()
	srv.Init()

	client := pb.NewUserService("user", srv.Client())
	rsp, err := client.Create(context.TODO(), &pb.User{
		Name:     "ZhangSan",
		Email:    "socylx@163.com",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("创建用户失败: %v", err)
	}
	log.Printf("创建用户成功: %s", rsp.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("获取所有用户失败: %v", err)
	}
	log.Println("err: ", err, getAll.Users)
	for _, v := range getAll.Users {
		log.Println(v)
	}

	time.Sleep(time.Second * 5)
}
