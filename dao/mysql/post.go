package mysql

import (
	"web_app/models"

	"go.uber.org/zap"
)

//CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id,title,content,author_id,community_id)
		values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

//GetPostById 根据帖子id获取帖子详情
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
				from post
				where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

//GetPostList 获取帖子列表
func GetPostList(page, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
				from post
				limit ?,?`

	post = make([]*models.Post, 0, 2)
	if err = db.Select(&post, sqlStr, (page-1)*size, size); err != nil {
		zap.L().Error("db.Select of GetPostList failed", zap.Error(err))
		return
	}
	return
}
