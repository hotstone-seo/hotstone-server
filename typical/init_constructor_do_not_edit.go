package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typdep"
)

func init() {
	typapp.AppendConstructor(
		typdep.NewConstructor(repository.NewDataSourceRepo),
		typdep.NewConstructor(repository.NewMetricsRuleMatchingRepo),
		typdep.NewConstructor(repository.NewRuleRepo),
		typdep.NewConstructor(repository.NewTagRepo),
		typdep.NewConstructor(repository.NewURLSyncRepo),
		typdep.NewConstructor(service.NewCenterService),
		typdep.NewConstructor(service.NewDataSourceService),
		typdep.NewConstructor(service.NewMetricsRuleMatchingService),
		typdep.NewConstructor(service.NewProviderService),
		typdep.NewConstructor(service.NewRuleService),
		typdep.NewConstructor(service.NewTagService),
		typdep.NewConstructor(service.NewURLService),
	)
}
