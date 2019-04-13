package conn

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"micro/ch04/config"
	"time"
)

var RedisDB *redis.Pool

func redisPool(maxIdle, maxActive int, idleTimeout time.Duration, network, address string, pwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, address)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _, err := c.Do("AUTH", pwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedis(c config.RedisConfig) {
	var maxIdle = 30 //连接池中最大空闲数
	var maxActive = 30 //连接池的最大数据库连接数
	var idleTimeout = 600 * time.Second
	var network = "tcp"
	var address = fmt.Sprintf("%s:%d", c.Address, c.Port)
	var pwd = ""
	RedisDB = redisPool(maxIdle, maxActive, idleTimeout, network, address, pwd)
	//使用的地方需要
	//conn := RedisDB.Get()
	//defer conn.Close()
}
