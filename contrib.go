package cache

import (
	"github.com/labstack/echo"
)

const EchoCacheStoreKey = "EchoCacheStore"

func Store(c interface{}) Cache {
	switch v := value.(type) {
	case *echo.Context:
		cacher := v.Get(EchoCacheStoreKey).(Cache)
		if cacher == nil {
			panic("EchoStore not found, forget to Use Middleware ?")
		}
	default:
		panic("unknown Context")
	}

	return cacher
}

func EchoCacher(opt Options) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			tagcache, err := New(opt)
			if err != nil {
				return err
			}

			c.Set(EchoCacheStoreKey, EchoCacheStore)

			h(c)
		}
	}
}
