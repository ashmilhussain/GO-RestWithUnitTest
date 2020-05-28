package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

/*****************************************************
	    model and crud operations for user
******************************************************/

type Post struct {
	PostID    int       `gorm:"primary_key;auto_increment"`
	PostName  string    `gorm:"size:40;not null;unique"`
	SubjectID int       `gorm:"not null"`
	IsDeleted bool      `gorm:"default:false" json:"_"`
	CreatedBy int       `json:"_"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
	UpdatedBy int       `json:"_"`
	UpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
}

func (post *Post) Prepare() {
	post.PostName = html.EscapeString(strings.TrimSpace(post.PostName))
}

func (post *Post) SavePost(db *gorm.DB) error {
	var err error
	post.Prepare()
	post.CreatedOn = time.Now()
	err = db.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (post *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	var posts []Post
	err = db.Where("is_deleted = ?", false).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	return &posts, err
}

func (post *Post) UpdatePost(db *gorm.DB) error {

	post.Prepare()
	post.UpdatedOn = time.Now()
	db = db.Model(Post{}).Where("post_id = ? and is_deleted = ? ", post.PostID, false).Update(&post)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (post *Post) DeletePost(db *gorm.DB) error {

	db = db.Model(Post{}).Where("post_id = ?", post.PostID).Updates(map[string]interface{}{"is_deleted": true})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
