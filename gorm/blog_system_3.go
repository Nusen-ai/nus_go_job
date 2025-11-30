package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目3：钩子函数
// 继续使用博客系统的模型。
// ● 要求 ：
//   ○ 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//   ○ 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

// Post 模型的钩子函数 - 在创建时更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("posts_count + ?", 1))

	if result.Error != nil {
		return result.Error
	}

	// 更新文章的评论状态
	result = tx.Model(&Post{}).Where("id = ?", p.ID).
		Update("comment_status", "无评论")

	return result.Error
}

// Post 模型的钩子函数 - 在删除时更新用户的文章数量
func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("GREATEST(posts_count - 1, 0)"))

	return result.Error
}

// Comment 模型的钩子函数 - 在创建时更新文章的评论数量和状态
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量
	result := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comments_count", gorm.Expr("comments_count + ?", 1))

	if result.Error != nil {
		return result.Error
	}

	// 更新文章的评论状态
	result = tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_status", "有评论")

	return result.Error
}

// Comment 模型的钩子函数 - 在删除时检查文章的评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 获取当前文章的评论数量
	var commentsCount int64
	result := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentsCount)
	if result.Error != nil {
		return result.Error
	}

	// 更新文章的评论数量
	result = tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comments_count", gorm.Expr("GREATEST(comments_count - 1, 0)"))

	if result.Error != nil {
		return result.Error
	}

	// 如果评论数量为0，更新评论状态为"无评论"
	if commentsCount == 0 {
		result = tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论")
	}

	return result.Error
}

// 完整的示例使用
func BlogSystem3TestDemo() {
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

	// 创建用户
	user := User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashed_password",
	}
	db.Create(&user)

	// 创建文章（会自动触发钩子函数更新用户文章数量）
	post := Post{
		Title:   "测试文章",
		Content: "这是测试文章的内容",
		UserID:  user.ID,
	}
	db.Create(&post)

	// 创建评论（会自动触发钩子函数更新文章评论数量和状态）
	comment := Comment{
		Content: "这是测试评论",
		UserID:  user.ID,
		PostID:  post.ID,
	}
	db.Create(&comment)

	// 查询示例
	// posts, _ := GetUserPostsWithComments(db, user.ID)
	// mostCommentedPost, _ := GetMostCommentedPost(db)

	// 删除评论（会自动触发钩子函数检查评论状态）
	db.Delete(&comment)
}
