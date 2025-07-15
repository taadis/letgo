package letgo

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/urfave/cli/v3"
)

type Application interface {
	Load(v any) error
	Run(ctx context.Context) error
}

type App struct {
	options *Options

	//
	kapp   *kratos.App
	cmd    *cli.Command
	logger log.Logger
}

func NewApp(opts ...Option) *App {
	options := NewOptions(opts...)

	a := new(App)
	a.options = options
	a.kapp = a.newkapp()
	a.cmd = a.newCmd()
	return a
}

func (a *App) Load(v any) error {
	return a.loadLocalConfig(v)
}

func (a *App) Run(ctx context.Context) error {
	return a.run(ctx)
}

func (a *App) run(ctx context.Context) error {
	return a.cmd.Run(ctx, os.Args)
}
