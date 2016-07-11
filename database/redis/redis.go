package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	p         *redis.Pool
	dbNum     int64
	conn      string
	password  string
	redisConn redis.Conn
}

func NewRedis() *Redis {
	return &Redis{}
}

func (this *Redis) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := this.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

func (this *Redis) Get(key string) interface{} {
	if v, err := this.do("GET", key); err == nil {
		return v
	}
	return nil
}

//func (this *Redis) GetMulti(keys []string) interface{} {
//	size := len(keys)
//	var rv []interface{}

//	c := this.p.Get()
//	defer c.Close()

//	var err error
//	for _, key := range keys {
//		err = c.Send("GET", key)
//		if err != nil {
//			break
//		}
//	}

//	if err != nil {
//		return nil
//	}

//	if err = c.Flush(); err != nil {
//		return nil

//	}

//	for i := 0; i < size; i++ {
//		if v, err := c.Receive(); err != nil {
//			rv = append(rv, v, []byte)
//		} else {
//			rv = append(rv, err)
//		}
//	}

//	return rv

//}

func (this *Redis) Put(key string, val interface{}, timeout int64) (err error) {
	if _, err = this.do("SETEX", key, timeout, val); err != nil {
		return
	}

	return
}

func (this *Redis) Delete(key string) (err error) {
	if _, err = this.do("DEL", key); err != nil {
		return
	}

	return
}

func (this *Redis) IsExist(key string) bool {
	v, err := redis.Bool(this.do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

//func (this *Redis) HSet(hash string, key string, value string) (err error) {
//	if _, err := this.do("HSET", hash, key, value); err != nil {
//		return
//	}
//	return
//}

func (this *Redis) HGet(hash string, key string) ([]byte, error) {
	value, err := redis.Bytes(this.do("HGET", hash, key))
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Redis) HGetAll(hash string) (map[string]int64, error) {
	value, err := redis.Int64Map(this.do("HGETALL", hash))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (this *Redis) HSet(hash string, key string, val interface{}) error {
	if _, err := this.do("HSET", hash, key, val); err != nil {
		return err
	}
	return nil
}

func (this *Redis) HIncrby(hash string, key string, val int64) (int64, error) {
	value, err := redis.Int64(this.do("HINCRBY", hash, key, val))
	if err != nil {
		return int64(0), err
	}
	return value, nil
}

func (this *Redis) HExist(hash string, key string) (bool, error) {
	value, err := redis.Bool(this.do("HEXIST", hash, key))
	if err != nil {
		return false, err
	}
	return value, nil
}

func (this *Redis) GetInt(key string) (int64, error) {
	value, err := redis.Int64(this.do("INCRBY", key, 0))
	if err != nil {
		return int64(0), err
	}
	return value, nil
}

func (this *Redis) Incr(key string) (int64, error) {
	value, err := redis.Int64(this.do("INCR", key))
	if err != nil {
		return int64(0), err
	}
	return value, nil
}

func (this *Redis) Expire(key string, second int64) error {
	_, err := this.do("EXPIRE", second)
	return err
}

func (this *Redis) Connect(conn string, dbNum int64, password string) (err error) {
	this.conn = conn
	this.dbNum = dbNum
	this.password = password

	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", this.conn)
		if err != nil {
			return nil, err
		}

		if this.password != "" {
			if _, err := c.Do("AUTH", this.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, err = c.Do("SELECT", this.dbNum)
		if err != nil {
			c.Close()
			return nil, err
		}

		this.redisConn = c
		return
	}

	this.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
	return
}
