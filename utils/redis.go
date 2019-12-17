package utils

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RailsCachePool *redis.Pool

	LimitPool *redis.Pool
	MiscPool  *redis.Pool
)

// 初始化各种连接池
func InitRedisPools() {
	RailsCachePool = newRedisPool("cache")
	LimitPool = newRedisPool("limit")

}

// 关闭连接池
func CloseRedisPools() {
	RailsCachePool.Close()
	LimitPool.Close()
}

// 根据别名获取连接池
func GetRedisConn(redisName string) redis.Conn {
	if redisName == "cache" {
		return RailsCachePool.Get()
	} else if redisName == "limit" {
		return LimitPool.Get()
	}
	return nil
}

// 创建redis连接池
func newRedisPool(redisName string) *redis.Pool {
	config := getRedisConfig()
	capacity := config.GetInt(redisName+"pool", 10)
	maxCapacity := config.GetInt(redisName+".maxopen", 0)
	idleTimeout := config.GetDuration(redisName+".timeout", "4m")
	maxConnLifetime := config.GetDuration(redisName+"life_time", "2m")
	network := config.Get(redisName+".network", "tcp")
	server := config.Get(redisName+".server", "localhost:6379")
	db := config.Get(redisName+".db", "")
	password := config.Get(redisName+".password", "")
	return &redis.Pool{
		MaxIdle:         capacity,
		MaxActive:       maxCapacity,
		IdleTimeout:     idleTimeout,
		MaxConnLifetime: maxConnLifetime,
		Wait:            true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(network, server)
			if err != nil {
				fmt.Println("无法连接redis: ", err.Error())
				return nil, err
			}
			if password != "" {
				_, err := conn.Do("AUTH", password)
				if err != nil {
					fmt.Println("redis认证失败: ", err.Error())
					conn.Close()
					return nil, err
				}
			}
			if db != "" {
				_, err := conn.Do("SELECT", db)
				if err != nil {
					fmt.Println("redis无法选择数据库: ", err.Error())
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			fmt.Println("开始检测redis连接")
			if err != nil {
				fmt.Println("redis无法检测ping: ", err.Error())
			}
			return err
		},
	}
}

type SubscribeCallback func(channel, message string)

type Subscriber struct {
	client redis.PubSubConn
	cbMap  map[string]SubscribeCallback
}

func Publish() {

}
