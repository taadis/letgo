package letgo

import (
	"github.com/go-kratos/kratos/v2"
)

func (a *App) newkapp() *kratos.App {
	opts := a.kopts()
	kapp := kratos.New(opts...)
	return kapp
}

func (a *App) kopts() []kratos.Option {
	var opts []kratos.Option
	if a.options.id != "" {
		opts = append(opts, kratos.ID(a.options.id))
	}
	if a.options.name != "" {
		opts = append(opts, kratos.Name(a.options.name))
	}
	if a.options.version != "" {
		opts = append(opts, kratos.Version(a.options.version))
	}
	if a.options.metadata != nil {
		opts = append(opts, kratos.Metadata(a.options.metadata))
	}
	if a.options.endpoints != nil {
		opts = append(opts, kratos.Endpoint(a.options.endpoints...))
	}
	if a.options.ctx != nil {
		opts = append(opts, kratos.Context(a.options.ctx))
	}
	if a.options.logger != nil {
		opts = append(opts, kratos.Logger(a.options.logger))
	}
	if a.options.servers != nil {
		opts = append(opts, kratos.Server(a.options.servers...))
	}
	if a.options.sigs != nil {
		opts = append(opts, kratos.Signal(a.options.sigs...))
	}
	if a.options.registrar != nil {
		opts = append(opts, kratos.Registrar(a.options.registrar))
	}
	if a.options.registrarTimeout != 0 {
		opts = append(opts, kratos.RegistrarTimeout(a.options.registrarTimeout))
	}
	if a.options.stopTimeout != 0 {
		opts = append(opts, kratos.StopTimeout(a.options.stopTimeout))
	}
	return opts
}
