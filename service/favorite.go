package service

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"strconv"
	"time"
)

func IsFavorite(ctx context.Context, userID int64, videoID int64) (isExist bool, isFavorite bool) {
	favoriteRelation, _ := model.QueryFavorite(ctx, userID, videoID)
	if favoriteRelation.ID == 0 {
		return false, false
	}
	if favoriteRelation.Status == 1 {
		return true, true
	}
	return true, false
}

func FavoriteAction(ctx context.Context, userID int64, videoID int64, actionType int) error {
	isExist, isFavorite := IsFavorite(ctx, userID, videoID)

	if !isExist {
		// 不存在的关系直接创建
		favorite := model.Favorite{
			UserId:  userID,
			VideoId: videoID,
			Status:  actionType,
		}
		if err := model.CreateFavorite(ctx, &favorite); err != nil {
			return err
		}
		dto.WriteLog("info", "点赞视频")
		// 重新刷新点赞列表
		LoadFavoriteListCache(ctx, userID)
	} else {
		// 存在的关系进行更新
		if (actionType == 1 && !isFavorite) || (actionType == 2 && isFavorite) {
			dto.WriteLog("info", "修改点赞状态")
			if err := model.UpdateFavorite(ctx, userID, videoID, actionType); err != nil {
				return err
			}
			// 重新刷新点赞列表
			LoadFavoriteListCache(ctx, userID)
		}
	}
	return nil
}

// load favorite video list cache
func LoadFavoriteListCache(ctx context.Context, userId int64) ([]dto.Video, error) {
	videoList, _, _ := model.QueryFavorites(ctx, userId, 0, 0)

	favoriteVideoList := make([]dto.Video, 0)
	for i := range videoList {
		video := dto.Video{}
		video.Id = videoList[i].ID
		video.PlayUrl = videoList[i].PlayUrl
		video.CoverUrl = videoList[i].CoverUrl
		video.FavoriteCount = videoList[i].FavoriteCount
		video.CommentCount = videoList[i].CommentCount
		author, _ := model.QueryUserById(ctx, videoList[i].AuthorID)
		user := dto.User{}
		user.Id = author.ID
		user.FollowerCount = author.FollowCount
		user.Name = author.Name
		user.FollowerCount = author.FollowerCount
		isExist, isFollow := IsFollow(ctx, userId, author.ID)
		if isExist && isFollow {
			user.IsFollow = true
		}
		video.Author = user
		video.IsFavorite = true
		favoriteVideoList = append(favoriteVideoList, video)
	}
	db.CacheSetList(context.Background(), "default", "favorite_video_list"+strconv.FormatInt(userId, 10), favoriteVideoList, time.Hour)
	dto.WriteLog("info", "刷新点赞列表的缓存")
	return favoriteVideoList, nil
}

func GetFavoriteList(ctx context.Context, userId int64) ([]dto.Video, error) {

	favoriteList, _ := db.CacheGetList(context.Background(), "default", "favorite_video_list"+strconv.FormatInt(userId, 10), []dto.Video{})
	if favoriteList != nil {
		dto.WriteLog("info", "获取点赞列表")
		return favoriteList, nil
	} else {
		favoriteList, _ = LoadFavoriteListCache(ctx, userId)
	}

	return favoriteList, nil
}
