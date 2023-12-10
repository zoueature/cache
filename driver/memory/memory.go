package memory

import (
	"context"
	"github.com/zoueature/cache"
	"github.com/zoueature/config"
	"sync"
	"time"
)

const DriverName = "memory"

func init() {
	cache.RegisterDriver(DriverName, func(conf config.CacheConfig) cache.Cache {
		return &ins{}
	})
}

type ins struct {
	data sync.Map
}

// TODO implement the memory cache

func (i *ins) Get(ctx context.Context, key string) string {
	//TODO implement me
	panic("implement me")
}

func (i *ins) Set(ctx context.Context, key string, value interface{}, ttl ...time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (i *ins) HSet(ctx context.Context, key, field string, value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (i *ins) HDelete(ctx context.Context, key string, field ...string) error {
	//TODO implement me
	panic("implement me")
}

func (i *ins) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
