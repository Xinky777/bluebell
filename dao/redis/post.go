package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

//getIDsFromKey 确定查询起始点 并排序
func getIDsFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + page - 1
	//3.ZRevRange 按分数从大到小查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

//GetPostIDsIOrder 查询id列表
func GetPostIDsIOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取id
	//1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderTime {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引起始点
	return getIDsFromKey(key, p.Page, p.Size)
}

//GetPostVoteData 根据ids获取帖子投票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//每个id都会发送redis请求 麻烦繁琐
	//data = make([]int64,0,len(ids))
	//for _,id := range ids{
	//	key := getRedisKey(KeyPostVotedZSetPF+id)
	//	//查找k中分数是1的元素数量 -> 统计每篇帖子的赞成票数量
	//	v := client.ZCount(ctx,key,"1","1").Val()
	//	data = append(data,v)
	//}

	//pipeline模式 一发送多条命令 减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

//GetCommunityPostIDsIOrder 按社区根据ids查询每篇帖子投赞成票的数据
func GetCommunityPostIDsIOrder(p *models.ParamPostList) ([]string, error) {
	//1.根据用户请求中携带的order参数确定要查询的redis key
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderTime {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	//社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	//使用zinTerStore 把分区的帖子set与帖子分数zset 生成一个新zset 针对新zset 按之前的逻辑取数据
	//利用缓存key减少zinTerStore执行次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(orderKey).Val() < 1 {
		//不存在 需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	//存在的话 根据key查询ids
	return getIDsFromKey(key, p.Page, p.Size)
}
