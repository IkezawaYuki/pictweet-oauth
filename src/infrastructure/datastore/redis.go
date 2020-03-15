package datastore

import (
	"github.com/gomodule/redigo/redis"
)

// RedisHandler is ...
type redisHandler struct {
	handler redis.Conn
}

// RedisHandler is ...
type RedisHandler interface {
	Get(string) string
	SetWithExpire(string, string) error
}

// NewRedisHandler is
func NewRedisHandler(c redis.Conn) RedisHandler {
	return &redisHandler{
		handler: c,
	}
}

// GetConnection Redisへの接続
func GetConnection() redis.Conn {
	const Addr = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return c
}

// Get 値の取得
func (h *redisHandler) Get(key string) string {
	res, _ := redis.String(h.handler.Do("GET", key))
	return res
}

// SetWithExpire 有効期限付きで値の設定
func (h *redisHandler) SetWithExpire(token string, email string) error {
	_, err := h.handler.Do("SET", token, email, "NX", "EX", "30")
	if err != nil {
		return err
	}
	return nil
}
