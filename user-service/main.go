package main

import (
	"fmt"
	micro "github.com/micro/micro/v3/service"
	"github.com/nats-io/nats.go"
	"github.com/socylx/laracom/user-service/model"
	"github.com/socylx/laracom/user-service/service"
	"log"
	"os"

	database "github.com/socylx/laracom/user-service/db"
	"github.com/socylx/laracom/user-service/handler"
	pb "github.com/socylx/laracom/user-service/proto/user"
	repository "github.com/socylx/laracom/user-service/repo"
)

func main() {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("Could not connect to NATS: %v", err)
	}

	// 创建数据库连接，程序退出时断开连接
	db, err := database.CreateConnection()
	defer func() { _ = db.Close() }()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// 和 Laravel 数据库迁移类似
	// 每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.PasswordReset{})

	// 初始化 Repo 实例用于后续数据库操作
	repo := &repository.UserRepository{Db: db.Debug()}
	resetRepo := &repository.PasswordResetRepository{Db: db}
	// 初始化 token service
	token := &service.TokenService{Repo: repo}

	// 以下是 Micro 创建微服务流程

	srv := micro.New(
		micro.Name("user"),
		micro.Version("latest"),
	)

	srv.Init()

	// 获取 Broker 实例
	//pubSub := srv.Server().Options().Broker

	// 注册处理器
	_ = pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{
		Repo:      repo,
		ResetRepo: resetRepo,
		Token:     token,
		PubSub:    nc,
	})

	// 启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
