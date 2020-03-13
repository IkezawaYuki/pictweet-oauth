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
	Get() (string, error)
	SetWithExpire(string, string) error
}

// NewRedisHandler is
func NewRedisHandler() RedisHandler {
	return &redisHandler{}
}

func (h *redisHandler) Get() (string, error) {

}

func (h *redisHandler) SetWithExpire(token string, email string) error {
	val, err := h.handler.Do("SET", email, token, "NX", "EX", "120")
	if err != nil {
		panic(err)
	}

	if val == nil {
		fmt.Println("already exist")
		os.Exit(1)
	}
}
