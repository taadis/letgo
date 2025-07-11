package letgo

import (
	"github.com/go-kratos/kratos/v2"
)

func (a *App) newkapp() *kratos.App {
	return kratos.New()
}
