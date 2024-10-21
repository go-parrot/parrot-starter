package repository

import (
	"github.com/go-parrot/parrot-starter/internal/model"
	"github.com/google/wire"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(model.Init, NewUserRepo)
