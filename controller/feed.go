package controller

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	token := c.Query("token")
	if token != "" {
		userInfo, parseTokenErr := util.ParseToken(token)
		if parseTokenErr != nil {
			dto.WriteLog(
				"error", "解析token错误",
			)
		}
		list, err := db.CacheGetList(context.Background(), "default", "feedList-"+strconv.FormatInt(userInfo.ID, 10), []dto.Video{})
		if err == nil {
			c.JSON(http.StatusOK, dto.FeedResponse{
				Response: dto.Response{StatusCode: 0,
					StatusMsg: "",
					NextTime:  time.Now().Unix()},
				VideoList: list,
			})
			return
		}
		videoList, reTime := service.QueryFeedResponse(userInfo.ID, lastTime)
		log.Println(reTime.Unix())
		err = db.CacheSetList(context.Background(), "default", "feedList-"+strconv.FormatInt(userInfo.ID, 10), videoList, time.Minute)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, dto.FeedResponse{
			Response: dto.Response{StatusCode: 0,
				StatusMsg: "",
				NextTime:  reTime.Unix()},
			VideoList: videoList,
		})
	} else {
		list, err := db.CacheGetList(context.Background(), "default", "noTokenFeedList", []dto.Video{})
		if err == nil {
			c.JSON(http.StatusOK, dto.FeedResponse{
				Response: dto.Response{StatusCode: 0,
					StatusMsg: "",
					NextTime:  time.Now().Unix()},
				VideoList: list,
			})
			return
		}
		videoList, reTime := service.NoTokenQueryFeedResponse(lastTime)
		err = db.CacheSetList(context.Background(), "default", "feedList", videoList, time.Minute)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, dto.FeedResponse{
			Response: dto.Response{StatusCode: 0,
				StatusMsg: "",
				NextTime:  reTime.Unix()},
			VideoList: videoList,
		})
	}

}
