package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/hotstone-seo/hotstone-seo/internal/analyt"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/internal/provider"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func init() {
	typgo.Provide(
		&typgo.Constructor{Name: "", Fn: oauth2google.NewService},
		&typgo.Constructor{Name: "", Fn: analyt.NewClientKeyAnalytRepo},
		&typgo.Constructor{Name: "", Fn: analyt.NewReportRepo},
		&typgo.Constructor{Name: "", Fn: analyt.NewRuleMatchingRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewAuditTrailRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewClientKeyRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewDataSourceRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewRuleRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewSettingRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewStructuredDataRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewTagRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewUserRepo},
		&typgo.Constructor{Name: "", Fn: repository.NewUserRoleRepo},
		&typgo.Constructor{Name: "", Fn: service.NewAuditTrailService},
		&typgo.Constructor{Name: "", Fn: service.NewAuthService},
		&typgo.Constructor{Name: "", Fn: service.NewCenterService},
		&typgo.Constructor{Name: "", Fn: service.NewClientKeyService},
		&typgo.Constructor{Name: "", Fn: service.NewDataSourceService},
		&typgo.Constructor{Name: "", Fn: service.NewMetricService},
		&typgo.Constructor{Name: "", Fn: service.NewUserRoleSvc},
		&typgo.Constructor{Name: "", Fn: service.NewRuleService},
		&typgo.Constructor{Name: "", Fn: service.NewSettingSvc},
		&typgo.Constructor{Name: "", Fn: service.NewStructuredDataService},
		&typgo.Constructor{Name: "", Fn: service.NewTagService},
		&typgo.Constructor{Name: "", Fn: service.NewUserSvc},
		&typgo.Constructor{Name: "", Fn: infra.Connect},
		&typgo.Constructor{Name: "", Fn: provider.NewService},
		&typgo.Constructor{Name: "", Fn: urlstore.NewService},
		&typgo.Constructor{Name: "", Fn: urlstore.NewStore},
		&typgo.Constructor{Name: "", Fn: urlstore.NewSyncRepo},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *infra.App, err error) {
				cfg = new(infra.App)
				if err = typgo.ProcessConfig("APP", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *infra.Redis, err error) {
				cfg = new(infra.Redis)
				if err = typgo.ProcessConfig("REDIS", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *infra.Pg, err error) {
				cfg = new(infra.Pg)
				if err = typgo.ProcessConfig("PG", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *infra.Analyt, err error) {
				cfg = new(infra.Analyt)
				if err = typgo.ProcessConfig("ANALYT", cfg); err != nil {
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
	)
	typgo.Destroy(
		&typgo.Destructor{Fn: infra.Disconnect},
	)
}
