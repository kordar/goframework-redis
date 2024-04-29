package goframework_redis

import (
	"github.com/go-redis/redis"
	"github.com/kordar/godb"
	log "github.com/kordar/gologger"
)

var redispool *godb.DbConnPool

func GetRedisClient(db string) *redis.Client {
	return redispool.Handle(db).(*redis.Client)
}

// InitRedisHandle 初始化redis句柄
func InitRedisHandle(dbs ...string) {
	for _, db := range dbs {
		ins := NewRedisConnIns(db)
		if ins == nil {
			continue
		}
		err := redispool.Add(ins)
		if err != nil {
			log.Warnf("初始化Redis异常，err=%v", err)
		}
	}
}

// AddRedisInstance 添加redis句柄
func AddRedisInstance(db string) error {
	ins := NewRedisConnIns(db)
	return redispool.Add(ins)
}

// RemoveRedisInstance 移除redis句柄
func RemoveRedisInstance(db string) {
	redispool.Remove(db)
}
