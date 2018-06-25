package help

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var Redis *RedisModel

func init() {
	var err error
	Redis, err = NewRedis()
	Error(err)
}

func NewRedis() (rc *RedisModel, err error) {
	host := beego.AppConfig.String("redis.host")
	port := beego.AppConfig.String("redis.port")
	if host == "" || port == "" {
		return nil, errors.New("配置错误")
	}

	rc = new(RedisModel)

	dialFunc := func() (c redis.Conn, err error) {
		//        c, err = redis.Dial("tcp", rc.conninfo)
		c, err = redis.Dial("tcp", host+":"+port)
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}

	return rc, nil
}

type RedisModel struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	dbNum    int
	key      string
	password string
}

func (rc *RedisModel) Set(key, val string) error {
	c := rc.p.Get()
	defer c.Close()

	_, err := c.Do("SET", key, val)
	Error(err)

	return err
}

func (rc *RedisModel) Get(key string) string {
	c := rc.p.Get()
	defer c.Close()
	val, err := redis.String(c.Do("GET", key))

	Error(err)

	return val
}

func (rc *RedisModel) Lpush(key, val string) error {
	c := rc.p.Get()
	defer c.Close()
	_, err := c.Do("lpush", key, val)

	Error(err)
	return err
}

func (rc *RedisModel) Rpop(key string) string {
	c := rc.p.Get()
	defer c.Close()
	val, err := redis.String(c.Do("rpop", key))

	Error(err)
	return val

}
