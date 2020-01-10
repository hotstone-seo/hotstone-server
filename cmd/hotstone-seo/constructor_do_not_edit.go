package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/hotstone-seo/hotstone-seo/app/metric"
	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"github.com/hotstone-seo/hotstone-seo/app/service"
	"github.com/hotstone-seo/hotstone-seo/app/urlstore"
	"github.com/hotstone-seo/hotstone-seo/typical"
)

func init() {
	typical.Descriptor.Constructors.Append(
		metric.NewMetricServer,
		repository.NewDataSourceRepo,
		repository.NewLocaleRepo,
		repository.NewMetricsRuleMatchingRepo,
		repository.NewRuleRepo,
		repository.NewTagRepo,
		service.NewCenterService,
		service.NewDataSourceService,
		service.NewLocaleService,
		service.NewMetricsRuleMatchingService,
		service.NewProviderService,
		service.NewRuleService,
		service.NewTagService,
		urlstore.NewURLStoreServer,
		urlstore.NewURLStoreSyncRepo,
		urlstore.NewURLStoreSyncService,
	)
}
