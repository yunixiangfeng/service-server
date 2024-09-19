package setup

import (
	"context"
	"fmt"
	// "git.internal.attains.cn/attains-cloud/service-acs/core/config/pkg"
	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context, driver *pkg.Redis) (redis.UniversalClient, error) {
	opt := &redis.UniversalOptions{
		Addrs:                 driver.Addrs,
		ClientName:            "",
		DB:                    driver.DB,
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              driver.Username,
		Password:              driver.Password,
		SentinelUsername:      "",
		SentinelPassword:      driver.SentinelPassword,
		MaxRetries:            driver.MaxRetries,
		ContextTimeoutEnabled: false,
		PoolFIFO:              driver.PoolFIFO,
		PoolSize:              driver.PoolSize,
		MinIdleConns:          driver.MinIdleConns,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		MaxRedirects:          driver.MaxRedirects,
		ReadOnly:              driver.ReadOnly,
		RouteByLatency:        driver.RouteByLatency,
		RouteRandomly:         driver.RouteRandomly,
		MasterName:            driver.MasterName,
	}
	redisClient := redis.NewUniversalClient(opt)
	if driver.Check {
		if err := redisClient.Ping(ctx).Err(); err != nil {
			return nil, fmt.Errorf("redis connect err: %v", err)
		}
	}
	return redisClient, nil
}
