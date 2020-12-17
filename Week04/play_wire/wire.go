package main

import (
	"context"
	"github.com/google/wire"
	"play_wire/foobarbaz"
)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}
