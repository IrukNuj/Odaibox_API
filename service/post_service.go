package post

import (
	"github.com/gin-gonic/gin"
	"github.com/iruknuj/odaibox_API/db"
	"github.com/iruknuj/odaibox_API/model"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
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

	config := &validator.Config{TagName: "v-post"}
	validate := validator.New(config)

	if err := validate.Struct(posts); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"body": "そういう事しないの"})
		return posts, err
	}

	if err := db.Create(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func (s Service) UpdateByID(id string, context *gin.Context) (Post, error) {
	db := db.GetDB()
	var post Post

	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		return post, err
	}

	if err := context.BindJSON(&post); err != nil {
		return post, err
	}

	db.Save(&post)

	return post, nil
}