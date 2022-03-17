package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

//投票的几种情况
/*
opection = 1：
	1.之前没有投过票 现在投赞成票  差值绝对值：1 +432
	2.之前投反对票 现在投赞成票    差值绝对值：2 +432*2

opection = 0：
	1.之前投赞成票 现在取消投票    差值绝对值：1 -432
	2.之前投反对票 现在取消投票    差值绝对值：1 +432

opection = -1：
	1.之前没有投过票 现在投反对票   差值绝对值：1 -432
	2.之前投赞成票 现在投反对票     差值绝对值：2 —432*2
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

//CreatePost 创建帖子投票时间的代码
func CreatePost(communityID, postID int64) error {
	//创建事物
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()), //帖子分数会随时间变化 越新帖子分数越高
		Member: postID,
	})
	//把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

//VoteForPost 为帖子投票代码
func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制
	//去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2和3需要放到一个pipeline中
	//2.更新帖子的分数
	//先查询之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val() //之前帖子投票分数 value为新的加分或减分
	var op float64                                                            //设定方向 1为加分 -1为减分
	//如果这一次投票的值和之前保存的值一致 提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	if ov > value {
		op = -1 //减分
	} else {
		op = 1 //加分
	}
	diff := math.Abs(ov - value) //计算两次投票的差值的绝对值

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	//3.记录用户为该帖子投票的数据
	if value == 0 { //取消投票
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
