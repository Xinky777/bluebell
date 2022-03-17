package models

//定义用户请求的参数结构体
//用户输入的结构体字段

//定义根据时间或分数排序的常量
const (
	OrderTime  = "time"
	OrderScore = "score"
)

//ParamSignUp 定义注册请求（Signup）的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//ParamLogin 定义登陆请求（Login）的参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ParamVoteData 定义投票数据的参数结构体
type ParamVoteData struct {
	//UserID
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票（1），反对票（-1）,取消投票（0）
}

//ParamPostList 获取帖子列表Query string 参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" from:"community_id"` //可以为空
	Page        int64  `json:"page" from:"page"`
	Size        int64  `json:"size" from:"size"`
	Order       string `json:"order" from:"order"`
}
