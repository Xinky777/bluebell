package redis

import (
	"context"
	"fmt"
	"time"
	"web_app/settings"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
)

//Init 初始化redis连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		//Password: viper.GetString("redis.host"),   // no password set 取不到默认为空
		DB:       cfg.Db,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
