package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = "localhost:30379"
	redisPass = "123456"
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,                // 空闲连接
		MaxActive:   30,                // 最大连接
		IdleTimeout: 300 * time.Second, // 空闲时间

		Dial: func() (conn redis.Conn, err error) {
			// 打开连接
			conn, err = redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println("Fail to dial tcp connect,error:", err.Error())
				return nil, err
			}
			// 认证
			if _, err = conn.Do("AUTH", redisPass); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},

		// 检测是否超时
		TestOnBorrow: func(conn redis.Conn, t time.Time) (err error) {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err = conn.Do("Ping")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
