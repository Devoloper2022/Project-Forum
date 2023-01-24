package dto

import "forum/internal/models"

type PostDto struct {
	ID        int64         `json:"ID"`
	Title     string        `json:"title"`
	Text      string        `json:"text"`
	Date      string        `json:"data"`
	User      UserDto       `json:"userId"`
	Like      int64         `json:"like"`
	Dislike   int64         `json:"dislike"`
	Likes     []PostLikeDto `json:"likes"`
	Categorys []CategoryDto `json:"categorys"`
	// Comments  []Comment     `json:"comments"`
}

func (p *PostDto) GetPostDto(post models.Post, user UserDto,[]PostLikeDto,catCategoryDto) PostDto {
	return PostDto{
		ID:      post.ID,
		Title:   post.Title,
		Text:    post.Text,
		Date:    post.Date.Format("d MMM yyyy HH:mm:ss"),
		Like:    post.Like,
		Dislike: post.Dislike,
	}
}

// func GetModel(post PostDto) *models.Post {
// 	date, _ := time.Parse("d MMM yyyy HH:mm:ss", post.Date)
// 	return &models.Post{
// 		ID:     post.ID,
// 		Title:  post.Title,
// 		Text:   post.Text,
// 		Date:   date,
// 		UserID: post.ID,
// 		Like:   post.Like,
// 	}
// }