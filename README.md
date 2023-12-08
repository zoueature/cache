# Cache

### 缓存包

### Quick Start

1. import
> import "gitlab.jiebu.com/base/cache/driver/redis"

2. New client
```go
import (
    "gitlab.jiebu.com/base/cache/driver/redis"
	"gitlab.jiebu.com/base/cache"
)

func main()  {
	cfg := &config.CacheConfig{}
	c := cache.New(redis.DriverName, cfg)
	world := c.Get("hello")
	println(world)
}


```