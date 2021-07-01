package middlewares

import (
	"context"
	"github.com/spf13/viper"
	"strconv"
	"time"
	"yurikoto.com/yurikoto-telegram-bot/redis"
)

// RateLimit频次控制
func RateLimit(sender string) bool {
	rdb := redis.GetRedis()
	ctx := context.Background()
	key := "yurikoto:telegram:rate_limit:" + sender
	limit := viper.GetInt("telegram.rate_limit.limit")
	ttl := viper.GetInt("telegram.rate_limit.ttl")

	exists, _ := rdb.Exists(ctx, key).Result()
	rdb.Incr(ctx, key)
	if exists != 0 {
		strRequestCnt, _ := rdb.Get(ctx, key).Result()
		requestCnt, _ := strconv.Atoi(strRequestCnt)
		if requestCnt > limit {
			return false
		}
	} else {
		rdb.Expire(ctx, key, time.Duration(ttl)*time.Second)
	}
	return true
}
