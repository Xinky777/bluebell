package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

//CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	//1.生成postID
	p.ID = snowflake.GenID()
	//2.保存到数据库
	return mysql.CreatePost(p)
	//3.返回
}

//GetPostById 根据帖子id查询帖子
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	//查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}

	//根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("Author_id", post.AuthorID), zap.Error(err))
		return
	}

	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}

	//初始化指针结构体 并完成数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

//GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	//查询并组合我们接口想用的数据
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	//初始化data结构体
	data = make([]*models.ApiPostDetail, 0, len(posts))
	//for循环 查询
	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed", zap.Int64("Author_id", post.AuthorID), zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
