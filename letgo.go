package letgo

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/urfave/cli/v3"
)

type Application interface {
	Run(ctx context.Context) error
}

type App struct {
	kapp   *kratos.App
	cmd    *cli.Command
	logger log.Logger
}

func NewApp() *App {
	a := new(App)
	a.kapp = a.newkapp()
	a.cmd = a.newCmd()
	return a
}

func (a *App) Run(ctx context.Context) error {
	return a.run(ctx)
}

func (a *App) run(ctx context.Context) error {
	return a.cmd.Run(ctx, os.Args)
}
