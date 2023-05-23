package main

import (
	"context"
	"echoDemo/ent"
	"echoDemo/user/controller"
	"echoDemo/user/repository"
	"echoDemo/user/usecase"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	e := echo.New()
	//连接数据库
	client, err := ent.Open("postgres", os.Getenv("LOCAL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	//关闭
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	repo := repository.NewRepository(client)
	uc := usecase.NewUsecase(repo)
	ctrl := controller.NewController(uc)

	e.Static("/", "public")

	//新建用户请求
	e.POST("/user/create", ctrl.CreateUser)
	//文件上传
	e.POST("/file/update", ctrl.Upload)
	//端口为9091
	e.Logger.Fatal(e.Start(":9091"))

}
