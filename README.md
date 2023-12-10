# Cache

### 缓存包

### Quick Start

1. import
> import "github.com/zoueature/cache/driver/redis"

2. New client
```go
import (
    "github.com/zoueature/cache/driver/redis"
	"github.com/zoueature/cache"
)

func main()  {
	cfg := &config.CacheConfig{}
	c := cache.New(redis.DriverName, cfg)
	world := c.Get("hello")
	println(world)
}


```
