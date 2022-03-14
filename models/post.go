package models

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

//ApiPostDetail 帖子详情借口的结构体
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	*Post                                      //嵌入帖子结构体
	*CommunityDetail `json:"community_detail"` //嵌入社区信息
}
