package models

import "time"

//Community 访问社区
type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

//CommunityDetail 社区访问详情
type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"` //omitempty 如果为空 不展示
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
