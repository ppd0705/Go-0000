package foobarbaz

import (
	"context"
	"errors"
	"github.com/google/wire"
)

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	X int
}

func ProvideBar(foo Foo) Bar {
	return Bar{X: foo.X}
}

type Baz struct {
	X int
}

func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)



