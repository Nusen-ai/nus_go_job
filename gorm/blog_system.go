package gorm

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// ● 要求 ：
//   ○ 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//   ○ 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"size:50;not null;unique"`
	Email      string `gorm:"size:100;not null;unique"`
	Password   string `gorm:"size:255;not null"`
	PostsCount int    `gorm:"default:0"` // 用户文章数量统计
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Posts      []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:200;not null"`
	Content       string `gorm:"type:text;not null"`
	UserID        uint   `gorm:"not null"`
	CommentStatus string `gorm:"size:20;default:'无评论'"` // 评论状态
	CommentsCount int    `gorm:"default:0"`             // 评论数量统计
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	PostID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
	Post      Post `gorm:"foreignKey:PostID"`
}

func BlogSystemTestDemo() {
	dsn := "username:password@tcp(localhost:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("failed to migrate database")
	}
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// ● 要求 ：
//   ○ 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//   ○ 编写Go代码，使用Gorm查询评论数量最多的文章信息。

// 查询某个用户发布的所有文章及其对应的评论信息
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post

	err := db.Preload("User").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Order("comments.created_at DESC")
		}).
		Where("user_id = ?", userID).
		Order("posts.created_at DESC").
		Find(&posts).Error

	return posts, err
}

// 查询评论数量最多的文章信息
func GetMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post

	err := db.Preload("User").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Order("comments.created_at DESC")
		}).
		Order("comments_count DESC").
		First(&post).Error

	return post, err
}

// 使用示例
func QueryExamples(db *gorm.DB) {
	// 查询用户ID为1的所有文章及评论
	posts, err := GetUserPostsWithComments(db, 1)
	if err != nil {
		// 处理错误
	}

	// 查询评论数量最多的文章
	mostCommentedPost, err := GetMostCommentedPost(db)
	if err != nil {
		// 处理错误
	}

	_ = posts
	_ = mostCommentedPost
}
