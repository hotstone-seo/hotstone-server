package typical

import (
	"github.com/hotstone-seo/hotstone-server/app"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/typical-go/typical-rest-server/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

// Descriptor of hotstone-server
var Descriptor = &typcore.ProjectDescriptor{
	Name:      "hotstone-server",
	Version:   "0.0.1",
	Package:   "github.com/hotstone-seo/hotstone-server",
	AppModule: app.Module(),
	Modules: []interface{}{
		typdocker.Module(),
		typserver.Module(),
		typredis.Module(),
		typpostgres.Module("hotstone"),
	},
	// ReadmeGenerator: typreadme.Generator{},
	Releaser: &typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
	},
}