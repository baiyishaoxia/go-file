package common

import (
	"context"
	"github.com/go-redis/redis/v8"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var RDB *redis.Client
var RedisEnabled = true

// InitRedisClient This function is called after init()
func InitRedisClient() (err error) {
	if appConfigInfo == nil || appConfigInfo.Redis == nil {
		RedisEnabled = false
		// The cache depends on Redis
		ExplorerCacheEnabled = false
		// This stat feature also depends on Redis
		StatEnabled = false
		return nil
	}
	//
	configInfo := appConfigInfo.Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     strings.Join([]string{configInfo.Host, ":", configInfo.Port}, ""),
		Password: configInfo.Password,
		PoolSize: configInfo.PoolSize,
		DB:       configInfo.DbNum,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//
	count := 0
	for {
		if count > 5 {
			break
		}
		if _, err := RDB.Ping(ctx).Result(); err != nil {
			_, file, line, _ := runtime.Caller(1)
			println(file+":"+strconv.Itoa(line), "Redis 连接测试：%s", err.Error())
			time.Sleep(1 * time.Second)
			count++
			continue
		}
		_, file, line, _ := runtime.Caller(0)
		println(file+":"+strconv.Itoa(line), "Redis %v", "连接成功")
		break
	}
	return err
}

func ParseRedisOption() *redis.Options {
	return RDB.Options()
}
