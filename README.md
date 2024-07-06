# goframework-redis

包装`redis`对象，实现[`godb`](https://github.com/kordar/godb)接口。

## 安装
```go
go get github.com/kordar/goframework-redis v1.0.8
```

## 使用

- 添加实例

```go
// 1、通过配置添加 
section := map[string]interface{}{
	"addr": "127.0.0.1", // Redis地址
	"password": "xxx",   // Redis账号
	"db": 12,   // Redis库
	"poolSize": 10,   // Redis连接池大小
	"maxRetries": 10,   // 最大重试次数
	"idleTimeout": 30,  // 空闲链接超时时间(秒)
	"minIdleConns": 3  // 空闲连接数量 	
}
if err := goframeworkredis.AddRedisInstance(key, section); err != nil {
    log.Errorf("初始化redis异常，err=%v", err)
}

// 2、通过redis的option实现
var options redis.Options = ...
goframeworkredis.AddMysqlInstanceWithDsn(key, options)
```



