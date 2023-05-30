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

	// 调用用户认证服务
	var token *pb.Token
	token, err = client.Auth(context.TODO(), &pb.User{
		Email:    "socylx@163.com",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("用户登录失败: %v", err)
	}
	log.Printf("用户登录成功：%s", token.Token)

	// 调用用户验证服务
	token, err = client.ValidateToken(context.TODO(), token)
	if err != nil {
		log.Fatalf("用户认证失败: %v", err)
	}
	log.Printf("用户认证成功：%s", token.Valid)

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
