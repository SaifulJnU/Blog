package bootstrap

import (
	"github.com/saifuljnu/blog/pkg/config"
	"github.com/saifuljnu/blog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoute()

	routing.Serve()
}
