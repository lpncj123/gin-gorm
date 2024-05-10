package models

import "github.com/gomodule/redigo/redis"

// 链接redis

var REDIS *redis.Pool // redis连接池

// 初始化redis连接
func NewRedis() {
	REDIS = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.111.140:6379")
		},
	}
}
