# cache


## import
```
	"github.com/weisd/cache"
	_ "github.com/weisd/cache/redis"
```

## echo Middleware
```
e := echo.New()
e.Use(cache.EchoCacher(cache.Options{Adapter: "redis", AdapterConfig: `{"Addr":":6379"}`, Section: "test", Interval: 5}))

e.Get("/test/cache/put", func(c *echo.Context) error {
	err := cache.Store(c).Put("name", "weisd", 10)
	if err != nil {
		return err
	}

	return c.String(200, "store ok")
})

e.Get("/test/cache/get", func(c *echo.Context) error {
	name := cache.Store(c).Get("name")

	return c.String(200, "get name %s", name)
})

e.Run(":1323")

```

##