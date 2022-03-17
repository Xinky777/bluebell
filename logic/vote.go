package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

//投票功能
//投一票加10分 86400/200 -> 200张赞成票可以给你的帖子续一天

//投票的几种情况
/*
direction = 1：
	1.之前没有投过票 现在投赞成票
	2.之前投反对票 现在投赞成票

direction = 0：
	1.之前投赞成票 现在取消投票
	2.之前投反对票 现在取消投票

direction = -1：
	1.之前没有投过票 现在投反对票
	2.之前投赞成票 现在投反对票

投票限制：
每个帖子发表之日起 一个星期之内允许用户投票 超过一个星期就不允许再投票了
	1.到期后将redis中保存的票数保存到mysql
	2.到期之后删除那个keyPostVotedZSetPF
*/

//VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost", //在终端debug错误
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
