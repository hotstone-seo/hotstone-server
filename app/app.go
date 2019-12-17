package app

import (
	"github.com/hotstone-seo/hotstone-server/app/config"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct{}

func (*module) Action() interface{} {
	return startServer
}

func (*module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "APP"
	spec = &config.Config{}
	loadFn = func(loader typcore.ConfigLoader) (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
