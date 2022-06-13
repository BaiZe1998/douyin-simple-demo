package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []dto.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	//需要指定最大上传尺寸，此处无法指定
	fmt.Println("进入publish")
	file, err := c.FormFile("data")
	if err != nil {
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")
	//上传视频,并添加一个video到数据库
	userIdFromC, _ := c.Get("user_id")
	userId, _ := userIdFromC.(int64)
	reVideo := service.UploadVideoAliyun(file, token, title, userId)
	//对两个redis缓存进行更改
	//publishList
	list, err := db.CacheGetList(context.Background(), "default", "publishList-"+strconv.FormatInt(userId, 10), []dto.Video{})
	if err == nil {
		video := dto.Video{
			Id:            reVideo.ID,
			Author:        dto.User{},
			PlayUrl:       reVideo.PlayUrl,
			CoverUrl:      reVideo.CoverUrl,
			FavoriteCount: reVideo.FavoriteCount,
			CommentCount:  reVideo.CommentCount,
			IsFavorite:    false,
			Title:         reVideo.Title,
		}
		videos := append([]dto.Video{video}, list...)
		db.CacheSetList(context.Background(), "default", "publishList-"+strconv.FormatInt(userId, 10), videos, time.Minute)
	}
	//feedList
	list, err = db.CacheGetList(context.Background(), "default", "feedList-"+strconv.FormatInt(userId, 10), []dto.Video{})
	if err == nil {
		id, err := model.QueryUserById(context.Background(), userId)
		if err != nil {
			return
		}
		video := dto.Video{
			Id: reVideo.ID,
			Author: dto.User{
				Id:            id.ID,
				Name:          id.Name,
				FollowCount:   id.FollowCount,
				FollowerCount: id.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       reVideo.PlayUrl,
			CoverUrl:      reVideo.CoverUrl,
			FavoriteCount: reVideo.FavoriteCount,
			CommentCount:  reVideo.CommentCount,
			IsFavorite:    false,
			Title:         reVideo.Title,
		}
		videos := append([]dto.Video{video}, list...)
		err = db.CacheSetList(context.Background(), "default", "feedList-"+strconv.FormatInt(userId, 10), videos, time.Minute)
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "test" + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	list, err := db.CacheGetList(context.Background(), "default", "publishList-"+c.Query("user_id"), []dto.Video{})
	if err == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: list,
		})
	} else {
		videoList := service.QueryPublishList1(userId)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
		err := db.CacheSetList(context.Background(), "default", "publishList-"+c.Query("user_id"), videoList, time.Minute)
		if err != nil {
			dto.WriteLog("error", "err", err.Error())
			return
		}
	}
}
