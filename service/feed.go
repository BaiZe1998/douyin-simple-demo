package service

import (
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"strconv"
	"time"
)

type videoListType struct {
	ID            int64
	AuthorId      int
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	CreatedAt     time.Time
	Title         string
	IsFavorite    int
	IsFollow      int

	AuthorName          string
	AuthorFollowCount   int64
	AuthorFollowerCount int64
}

func QueryFeedResponse(useId int64, lastTime string) ([]dto.Video, time.Time) {
	var videoList []dto.Video = make([]dto.Video, 10)
	//query video list for feed
	if lastTime == "0" {
		now := time.Now()
		lastTime = now.Format("2006-01-02 15:04:05")
	} else {
		parseInt, err := strconv.ParseInt(lastTime, 10, 64)
		if err != nil {
			return nil, time.Time{}
		}
		lastTime = time.Unix(parseInt, 0).Format("2006-01-02 15:04:05")
	}
	// _, res := model.QueryVideoList(context.Background(), lastTime)

	var res []videoListType
	sqlQuery := "SELECT video.*, IFNULL(favoriteList.status, 2) AS IsFavorite, IFNULL(followList.status, 2) as IsFollow, user.name AS AuthorName, user.follow_count AS AuthorFollowCount, user.follower_count AS AuthorFollowerCount FROM video LEFT JOIN (SELECT video_id, user_id, status FROM favorite WHERE user_id =  11) AS favoriteList ON video.id = favoriteList.video_id LEFT JOIN ( SELECT followed_user, status FROM follow WHERE user_id = 11) AS followList ON video.author_id = followList.followed_user LEFT JOIN user ON video.author_id=user.id ORDER BY id DESC LIMIT 10;"
	queryErr := model.DB.Raw(sqlQuery).Scan(&res).Error

	if queryErr != nil {
		fmt.Println(queryErr)
	}

	at := res[len(res)-1].CreatedAt

	for index, value := range res {

		var is_follow bool
		var is_fav bool

		if value.IsFollow == 1 {
			is_follow = true
		} else {
			is_follow = false
		}

		if value.IsFavorite == 1 {
			is_fav = true
		} else {
			is_fav = false
		}

		videoList[index] = dto.Video{
			Id: value.ID,
			Author: dto.User{
				Id:            int64(value.AuthorId),
				Name:          value.AuthorName,
				FollowCount:   value.AuthorFollowCount,
				FollowerCount: value.AuthorFollowerCount,
				IsFollow:      is_follow,
			},
			PlayUrl:       value.PlayUrl,
			CoverUrl:      value.CoverUrl,
			FavoriteCount: value.FavoriteCount,
			CommentCount:  value.CommentCount,
			IsFavorite:    is_fav,
			Title:         value.Title,
		}
	}
	//return videoList[0:len(res)], at
	return videoList[0:len(res)], at
}

func NoTokenQueryFeedResponse(lastTime string) ([]dto.Video, time.Time) {
	var videoList []dto.Video = make([]dto.Video, 10)
	//query video list for feed
	if lastTime == "0" {
		now := time.Now()
		lastTime = now.Format("2006-01-02 15:04:05")
	} else {
		parseInt, err := strconv.ParseInt(lastTime, 10, 64)
		if err != nil {
			return nil, time.Time{}
		}
		lastTime = time.Unix(parseInt, 0).Format("2006-01-02 15:04:05")
	}
	// _, res := model.QueryVideoList(context.Background(), lastTime)

	var res []videoListType
	sqlQuery := "SELECT video.*,user.name AS AuthorName,user.follow_count AS AuthorFollowCount,user.follower_count AS AuthorFollowerCount FROM video , user WHERE video.author_id=user.id ORDER BY created_at DESC LIMIT 10 ;"
	queryErr := model.DB.Raw(sqlQuery).Scan(&res).Error

	if queryErr != nil {
		fmt.Println(queryErr)
	}

	at := res[len(res)-1].CreatedAt

	for index, value := range res {

		var is_follow bool
		var is_fav bool

		if value.IsFollow == 1 {
			is_follow = true
		} else {
			is_follow = false
		}

		if value.IsFavorite == 1 {
			is_fav = true
		} else {
			is_fav = false
		}

		videoList[index] = dto.Video{
			Id: value.ID,
			Author: dto.User{
				Id:            int64(value.AuthorId),
				Name:          value.AuthorName,
				FollowCount:   value.AuthorFollowCount,
				FollowerCount: value.AuthorFollowerCount,
				IsFollow:      is_follow,
			},
			PlayUrl:       value.PlayUrl,
			CoverUrl:      value.CoverUrl,
			FavoriteCount: value.FavoriteCount,
			CommentCount:  value.CommentCount,
			IsFavorite:    is_fav,
			Title:         value.Title,
		}
	}
	//return videoList[0:len(res)], at
	return videoList[0:len(res)], at
}
