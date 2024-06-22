package goframework_redis

import (
	"github.com/go-redis/redis"
	log "github.com/kordar/gologger"
	"github.com/spf13/cast"
	"time"
)

type RedisConnIns struct {
	name string
	ins  *redis.Client
}

func NewRedisConnIns(name string, cfg map[string]string) *RedisConnIns {
	options := redis.Options{
		Addr:         cfg["addr"],                                       // Redis地址
		Password:     cfg["password"],                                   // Redis账号
		DB:           cast.ToInt(cfg["db"]),                             // Redis库
		PoolSize:     cast.ToInt(cfg["poolSize"]),                       // Redis连接池大小
		MaxRetries:   cast.ToInt(cfg["maxRetries"]),                     // 最大重试次数
		IdleTimeout:  cast.ToDuration(cfg["idleTimeout"]) * time.Second, // 空闲链接超时时间
		MinIdleConns: cast.ToInt(cfg["minIdleConns"]),                   // 空闲连接数量
	}
	return NewRedisConnInsWithRedisOption(name, options)
}

func NewRedisConnInsWithRedisOption(name string, option redis.Options) *RedisConnIns {
	client := redis.NewClient(&option)
	conn := RedisConnIns{name: name}
	if ok := conn.Ping(client); ok {
		conn.ins = client
		return &conn
	} else {
		log.Error("实例化redis client异常")
		return nil
	}
}

func (c RedisConnIns) GetName() string {
	return c.name
}

func (c RedisConnIns) GetInstance() interface{} {
	return c.ins
}

func (c RedisConnIns) Ping(client *redis.Client) bool {
	pong, err := client.Ping().Result()
	if err == redis.Nil {
		log.Warn("Redis异常")
		return false
	} else if err != nil {
		log.Warn("失败:", err)
		return false
	} else {
		log.Info(pong)
		return true
	}
}

func (c RedisConnIns) Close() error {
	return c.ins.Close()
}
