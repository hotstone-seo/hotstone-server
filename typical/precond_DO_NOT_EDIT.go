package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/hotstone-seo/hotstone-seo/server/infra"
	"github.com/hotstone-seo/hotstone-seo/server/metric"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/hotstone-seo/hotstone-seo/server/urlstore"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
)

func init() {
	typgo.Provide(
		&typgo.Constructor{Name: "", Fn: infra.Connect},
		&typgo.Constructor{Name: "", Fn: metric.NewReportRepo},
		&typgo.Constructor{Name: "", Fn: metric.NewRuleMatchingRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewAuditTrailRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewClientKeyRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewDataSourceRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewHistoryRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewRoleTypeRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewRuleRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewStructuredDataRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewTagRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewURLSyncRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewUserRepo},
		&typgo.Constructor{Name: "", Fn: service.NewAuditTrailService},
		&typgo.Constructor{Name: "", Fn: service.NewCenterService},
		&typgo.Constructor{Name: "", Fn: service.NewClientKeyService},
		&typgo.Constructor{Name: "", Fn: service.NewDataSourceService},
		&typgo.Constructor{Name: "", Fn: service.NewHistoryService},
		&typgo.Constructor{Name: "", Fn: service.NewMetricService},
		&typgo.Constructor{Name: "", Fn: service.NewProviderService},
		&typgo.Constructor{Name: "", Fn: service.NewRoleTypeService},
		&typgo.Constructor{Name: "", Fn: service.NewRuleService},
		&typgo.Constructor{Name: "", Fn: service.NewStructuredDataService},
		&typgo.Constructor{Name: "", Fn: service.NewTagService},
		&typgo.Constructor{Name: "", Fn: service.NewURLService},
		&typgo.Constructor{Name: "", Fn: service.NewUserService},
		&typgo.Constructor{Name: "", Fn: urlstore.NewStore},
		&typgo.Constructor{Name: "", Fn: oauth2google.NewService},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *config.Config, err error) {
				cfg = new(config.Config)
				if err = typgo.ProcessConfig("APP", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *oauth2google.Config, err error) {
				cfg = new(oauth2google.Config)
				if err = typgo.ProcessConfig("OAUTH2_GOOGLE", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *typredis.Config, err error) {
				cfg = new(typredis.Config)
				if err = typgo.ProcessConfig("REDIS", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *typpostgres.Config, err error) {
				cfg = new(typpostgres.Config)
				if err = typgo.ProcessConfig("PG", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
	)
	typgo.Destroy(
		&typgo.Destructor{Fn: infra.Disconnect},
	)
}
