package main

import (
	user "github.com/Jeff634-2/user/web/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
)

//type RegisterParameters struct {
//	UserName  string `form:"UserName" json:"UserName"`
//	FirstName string `form:"FirstName" json:"FirstName"`
//	Pwd       string `form:"Pwd" json:"Pwd"`
//}

func ServiceUser(c *gin.Context) {
	service := micro.NewService()

	service.Init()

	//创建微服务客户端
	client := user.NewUserService("go.micro.service.user", service.Client())

	//调用服务
	rsp, err := client.Register(c, &user.UserRegisterRequest{
		UserName:  c.Query("UserName"),
		FirstName: c.Query("FirstName"),
		Pwd:       c.Query("Pwd"),
	})

	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": rsp.Message})
}

func UserInfo(c *gin.Context) {
	service := micro.NewService()

	service.Init()

	//创建微服务客户端
	client := user.NewUserService("go.micro.service.user", service.Client())

	//调用服务
	rsp, err := client.GetUserInfo(c, &user.UserInfoRequest{
		UserName: c.Query("UserName"),
	})

	if err != nil {
		c.JSON(200, gin.H{"code": 300, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": rsp})
}

func main() {
	r := gin.Default()
	r.POST("userRegister", ServiceUser)
	r.GET("userInfo/:UserName", UserInfo)
	r.Run(":8080")
}
