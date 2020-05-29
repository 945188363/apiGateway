package RedisUtil

import (
	redisPool "apiGateway/Common/Catch/redis"
	"apiGateway/Utils/ComponentUtil"
	"github.com/garyburd/redigo/redis"
)

const (
	redisLockTimeout = 10 // 10 seconds
)

func TryLock(k string) (isLock bool, err error) {
	// 获取连接
	conn := redisPool.RedisPool().Get()
	// 关闭连接
	defer conn.Close()
	// 这里需要redis.String包一下，才能返回redis.ErrNil
	_, err = redis.String(conn.Do("SET", k, 1, "ex", redisLockTimeout, "nx"))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	isLock = true
	return
}

func Unlock(k string) (err error) {
	// 获取连接
	conn := redisPool.RedisPool().Get()
	// 关闭连接
	defer conn.Close()
	_, err = conn.Do("DEL", k)
	if err != nil {
		return
	}
	return
}

func Get(k string) (interface{}, error) {
	// 获取连接
	conn := redisPool.RedisPool().Get()
	defer conn.Close()
	v, err := conn.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 插入数据
func Set(k string, v interface{}) error {
	// 获取连接
	conn := redisPool.RedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("SET", k, v)
	return err
}

// 删除数据
func Del(k string) error {
	// 获取连接
	conn := redisPool.RedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("DEL", k)
	return err
}

// 插入数据，设置过期时间，秒
func SetEx(k string, v interface{}, ex int64) error {
	// 获取连接
	conn := redisPool.RedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("SET", k, v, "EX", ex)
	return err
}

// 查询是否存在
func Exist(k string) bool {
	// 获取连接RedisUtil.Get("Name")
	conn := redisPool.RedisPool().Get()
	defer conn.Close()
	exist, err := redis.Bool(conn.Do("EXISTS", k))
	if err != nil {
		ComponentUtil.RuntimeLog().Warn("redis util error,error:", err.Error())
		return false
	}
	return exist
}

// 设置过期时间，秒
func Expire(k string, times int) bool {
	// 获取连接RedisUtil.Get("Name")
	conn := redisPool.RedisPool().Get()
	defer conn.Close()
	exist, err := redis.Bool(conn.Do("EXPIRE", k, times))
	if err != nil {
		ComponentUtil.RuntimeLog().Warn("redis util error,error:", err.Error())
		return false
	}
	return exist
}
