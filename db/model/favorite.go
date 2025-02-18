package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	VideoId   int64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// video表点赞数+1
	if err := UpdateVideoFavorite(ctx, favorite.VideoId, 1); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("favorite").WithContext(ctx).Create(favorite).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func UpdateFavorite(ctx context.Context, userID, videoID int64, status int) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := DB.Table("favorite").Raw("UPDATE favorite SET status=%d WHERE user_id=%d and video_id=%d", status, userID, videoID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if status == 2 {
		status = -1
	}
	if err := UpdateVideoFavorite(ctx, videoID, status); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func QueryFavorite(ctx context.Context, userID int64, videoID int64) (Favorite, error) {
	var favorite Favorite

	if err := DB.Table("favorite").WithContext(ctx).Where("user_id=? and video_id=? ", userID, videoID).First(&favorite).Error; err != nil {
		return Favorite{}, nil
	}
	return favorite, nil
}

func QueryFavorites(ctx context.Context, userID int64, limit, offset int) ([]Video, int64, error) {
	var videoList []Video
	var total int64
	var conn *gorm.DB
	conn = DB.Table("favorite").WithContext(ctx).Joins("inner join video on favorite.video_id = video.id").
		Select("video.id", "video.author_id", "video.play_url", "video.cover_url", "video.favorite_count", "video.comment_count").
		Where("favorite.user_id = ?", userID)

	if err := conn.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit == 0 {
		if err := conn.Scan(&videoList).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := conn.Limit(limit).Offset(offset).Scan(&videoList).Error; err != nil {
			return nil, 0, err
		}
	}
	return videoList, total, nil
}
