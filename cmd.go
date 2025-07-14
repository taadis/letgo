package letgo

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"
)

func (a *App) newCmd() *cli.Command {
	return &cli.Command{
		Name:    a.options.name,
		Version: a.options.version,
		Flags:   a.newCommandFlags(),
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if a.kapp == nil {
				a.kapp = a.newkapp()
			}

			err := a.kapp.Run()
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
}

func (a *App) newCommandFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Usage:   "指定运行环境 (如 dev, prod)",
			Value:   a.options.Env,
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "指定 HTTP 服务端口",
			Value:   a.options.Port,
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "指定配置文件路径",
			Value:   a.options.ConfigPath,
		},
	}
}
