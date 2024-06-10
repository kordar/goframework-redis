package goframework_redis

import (
	"errors"
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
		err := AddRedisInstance(db)
		if err != nil {
			log.Warnf("初始化Redis异常，err=%v", err)
		}
	}
}

// AddRedisInstance 添加redis句柄
func AddRedisInstance(db string) error {
	ins := NewRedisConnIns(db)
	if ins == nil {
		return errors.New("create conn fail")
	}
	return redispool.Add(ins)
}

// RemoveRedisInstance 移除redis句柄
func RemoveRedisInstance(db string) {
	redispool.Remove(db)
}

// HasRedisInstance redis句柄是否存在
func HasRedisInstance(db string) bool {
	return redispool != nil && redispool.Has(db)
}
