package letgo

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"
)

func (a *App) newCmd(opts ...Option) *cli.Command {
	options := NewOptions(opts...)

	return &cli.Command{
		Name:    options.Name,
		Version: options.Version,
		Flags:   a.newCommandFlags(options),
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

func (a *App) newCommandFlags(options *Options) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Usage:   "指定运行环境 (如 dev, prod)",
			Value:   options.Env,
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "指定 HTTP 服务端口",
			Value:   options.Port,
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "指定配置文件路径",
			Value:   options.ConfigPath,
		},
	}
}
