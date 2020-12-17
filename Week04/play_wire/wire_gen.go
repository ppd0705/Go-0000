// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"play_wire/foobarbaz"
)

// Injectors from wire.go:

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	foo := foobarbaz.ProvideFoo()
	bar := foobarbaz.ProvideBar(foo)
	baz, err := foobarbaz.ProvideBaz(ctx, bar)
	if err != nil {
		return foobarbaz.Baz{}, err
	}
	return baz, nil
}