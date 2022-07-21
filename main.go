package main

import (
	"fmt"
	"github.com/Jeff634-2/user/domain/repository"
	"github.com/Jeff634-2/user/handler"
	user "github.com/Jeff634-2/user/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	//注册consul

	// 服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	// 初始化服务
	srv.Init()

	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/micro?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("======>数据库连接成功<======")
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	// 一定要记得关闭
	defer db.Close()

	// 禁止创建附表名称
	db.SingularTable(true)

	// 创建服务实例
	//userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	// 注册handler
	err = user.RegisterUserHandler(srv.Server(), new(handler.User))

	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
