package repositories

import (
	"PockitGolangBoilerplate/models"

	"github.com/jinzhu/gorm"
)

type PostRepositoryQ interface {
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, id int)
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (postRepository *PostRepository) GetPosts(posts *[]models.Post) {
	postRepository.DB.Find(posts)
}

func (postRepository *PostRepository) GetPost(post *models.Post, id int) {
	postRepository.DB.Where("id = ? ", id).Find(post)
}
