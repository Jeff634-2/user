package handler

import (
	"context"
	"fmt"
	"github.com/Jeff634-2/user/domain/model"
	"github.com/Jeff634-2/user/domain/service"
	user "github.com/Jeff634-2/user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

// 注册
func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {

	//userRegister := &model.User{}
	////var userRegister model.User
	////if err := common.SwapTo(userRegisterRequest, userRegister); err != nil {
	////	return err
	////}
	//
	//dataByte, err := json.Marshal(userRegisterRequest)
	//if err != nil {
	//	return err
	//}
	//err = json.Unmarshal(dataByte, userRegister)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = u.UserDataService.AddUser(userRegister)
	//if err != nil {
	//	return err
	//}
	//userRegisterResponse.Message = "添加成功"
	//return nil

	var userRegister *model.User
	userRegister = new(model.User)
	//userRegister = make(model.User, 0)
	userRegister = &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	fmt.Println(userRegister)
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

// 登录
func (u *User) Login(ctx context.Context, userLogin *user.UserLoginRequest, loginResponse *user.UserLoginResponse) error {
	isOK, err := u.UserDataService.CheckPwd(userLogin.UserName, userLogin.Pwd)
	if err != nil {
		return err
	}
	loginResponse.IsSuccess = isOK
	return nil

}

// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = UserForResponse(userInfo)
	return nil
}

// 类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
