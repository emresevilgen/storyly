package inmemory_cache_service

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Service interface {
	Set(key string, val interface{}, duration time.Duration)
	Get(key string) (val interface{}, found bool)
}

type service struct {
	cache *cache.Cache
}

func New(c *cache.Cache) Service {
	return &service{cache: c}
}

func (s *service) Set(key string, val interface{}, duration time.Duration) {
	s.cache.Set(key, val, duration)
}

func (s *service) Get(key string) (val interface{}, found bool) {
	return s.cache.Get(key)
}
