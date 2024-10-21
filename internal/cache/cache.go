package cache

import (
	"github.com/go-parrot/parrot/pkg/redis"
	"github.com/google/wire"
)

// ProviderSet is cache providers.
var ProviderSet = wire.NewSet(redis.Init, NewUserCache)
