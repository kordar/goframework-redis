package goframework_redis

import (
	"github.com/go-redis/redis"
	"github.com/kordar/godb"
	logger "github.com/kordar/gologger"
)

var (
	redispool = godb.NewDbPool()
)

func GetRedisClient(db string) *redis.Client {
	return redispool.Handle(db).(*redis.Client)
}

// AddRedisInstances 批量添加redis句柄
func AddRedisInstances(dbs map[string]map[string]string) {
	for db, cfg := range dbs {
		ins := NewRedisConnIns(db, cfg)
		if ins == nil {
			continue
		}
		err := redispool.Add(ins)
		if err != nil {
			logger.Warnf("[godb-redis] 初始化Redis异常，err=%v", err)
		}
	}
}

// AddRedisInstance 添加redis句柄
func AddRedisInstance(db string, cfg map[string]string) error {
	ins := NewRedisConnIns(db, cfg)
	return redispool.Add(ins)
}

func AddRedisInstanceWithRedisOptions(db string, option redis.Options) error {
	ins := NewRedisConnInsWithRedisOption(db, option)
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
