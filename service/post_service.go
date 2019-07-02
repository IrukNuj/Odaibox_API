package post

import (
	"github.com/gin-gonic/gin"
	"github.com/iruknuj/odaibox_API/db"
	"github.com/iruknuj/odaibox_API/model"
)

type Service struct {
}

type Post model.Post

// 全postの取得
func (s Service) GetAll() ([]Post, error) {
	db := db.GetDB()
	var posts []Post

	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// Post modelの生成
func (s Service) CreateModel(context *gin.Context) (Post, error) {
	db := db.GetDB()
	var posts Post

	if err := context.BindJSON(&posts); err != nil {
		return posts, err
	}

	if err := db.Create(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}
