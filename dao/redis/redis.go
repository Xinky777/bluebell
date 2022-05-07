package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

//Init 初始化redis连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		//Password: viper.GetString("redis.host"),   // no password set 取不到默认为空
		DB:       cfg.Db,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
