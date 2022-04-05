package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"url-shortener/internal/service/config"
)

type NoSql struct {
	pool *redis.Pool
}

func NewNosql(c config.Config) NoSql {
	pool := &redis.Pool{
		MaxIdle:   c.NoSql.MaxIdle,
		MaxActive: c.NoSql.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(c.NoSql.Network, c.NoSql.Addr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return NoSql{pool: pool}
}

func (noSql NoSql) Set(key, value string, expireTime int) error {
	client := noSql.pool.Get()
	defer client.Close()

	_, err := client.Do("SETEX", key, expireTime, value)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (noSql NoSql) Get(key string) (string, error) {
	client := noSql.pool.Get()
	defer client.Close()

	value, err := client.Do("GET", key)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	if value == nil {
		return "", nil
	}
	url := fmt.Sprintf("%s", value)
	return url, nil
}

func (noSql NoSql) Close() error {
	return noSql.pool.Close()
}
