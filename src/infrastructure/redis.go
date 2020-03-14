package infrastructure

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

// RedisHandler is ...
type redisHandler struct {
	handler redis.Conn
}

// RedisHandler is ...
type RedisHandler interface {
	Get(string) (string, error)
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
func (h *redisHandler) Get(key string) (string, error) {
	res, err := redis.String(h.handler.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res, nil
}

// SetWithExpire 有効期限付きで値の設定
func (h *redisHandler) SetWithExpire(token string, email string) error {
	val, err := h.handler.Do("SET", email, token, "NX", "EX", "10")
	if err != nil {
		return err
	}

	if val == nil {
		fmt.Println("already exist")
		os.Exit(1)
	}

	return nil
}
