package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var userIdSequence = int64(1)

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//Password encrypted with salt
	password, _ = service.Encryption(password)

	//QueryUser QueryUser By Name for judged user is exit or not
	user, _ := model.QueryUserByName(context.Background(), username)

	//judege user exit or not
	if user.Name != "" {
		log.Printf(user.Name, user.ID)
		c.JSON(http.StatusOK, dto.UserLoginResponse{
			Response: dto.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := &model.User{
			Name:            username,
			Password:        password,
			BackgroundImage: "https://tse2-mm.cn.bing.net/th/id/OIP-C.sDoybxmH4DIpvO33-wQEPgHaEq?pid=ImgDet&rs=1",
			Avatar:          "https://p3-passport.byteacctimg.com/img/user-avatar/dbc6c60e44668ebc05e930c5c3c3e8e7~300x300.image",
			Signature:       "Go语言学习中...",
		}
		//userinfo register
		model.CreateUser(context.Background(), newUser)
		//Query Userinfo for get id
		userInfo, _ := model.QueryUserByName(context.Background(), username)
		//token
		token, _ := util.GenerateToken(&util.UserClaims{ID: userInfo.ID, Name: username, PassWord: password})
		c.JSON(http.StatusOK, dto.UserLoginResponse{
			Response: dto.Response{StatusCode: 0},
			UserId:   userInfo.ID,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//Password encrypted with salt
	encryptionPassWord, _ := service.Encryption(password)

	user, _ := model.QueryUserByName(context.Background(), username)
	token, _ := util.GenerateToken(&util.UserClaims{ID: user.ID, Name: username, PassWord: encryptionPassWord})

	if user != nil {
		//judge password
		if service.ComparePasswords(user.Password, password) {
			c.JSON(http.StatusOK, dto.UserLoginResponse{
				Response: dto.Response{StatusCode: 0},
				UserId:   user.ID,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, dto.UserLoginResponse{
				Response: dto.Response{StatusCode: 1, StatusMsg: "password wrong"},
			})
		}
	} else {
		c.JSON(http.StatusOK, dto.UserLoginResponse{
			Response: dto.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	userClaims, _ := util.ParseToken(token)
	userModel, _ := model.QueryUserById(context.Background(), userClaims.ID)

	user := dto.User{
		Id:              userModel.ID,
		Name:            userModel.Name,
		FollowCount:     userModel.FollowCount,
		FollowerCount:   userModel.FollowerCount,
		IsFollow:        false,
		BackgroundImage: userModel.BackgroundImage,
		Avatar:          userModel.Avatar,
		Signature:       userModel.Signature,
	}
	c.JSON(http.StatusOK, dto.UserResponse{
		Response: dto.Response{StatusCode: 0},
		User:     user,
	})

	//log.Printf(token)
	//log.Printf(userinfo.ID)
}
