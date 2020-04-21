package service

import (
	"context"
	"net/url"

	"github.com/go-redis/redis"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/metric"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for provider api [mock]
type ProviderService interface {
	Match(context.Context, MatchRequest) (*MatchResponse, error)
	FetchTags(context.Context, url.Values) ([]*ITag, error)
	FetchTagsWithCache(context.Context, url.Values, *cachekit.Pragma) ([]*ITag, error)
}

// ProviderServiceImpl is implementation of Provider
type ProviderServiceImpl struct {
	dig.In
	metric.RuleMatchingRepo
	repository.DataSourceRepo
	repository.RuleRepo
	repository.TagRepo
	URLService

	Redis *redis.Client
}

// ITag is tag after interpolate with data
type ITag repository.Tag

// IDataSource is datasource after interpolate with data
type IDataSource repository.DataSource

// NewProviderService return new instance of ProviderService [constructor]
func NewProviderService(impl ProviderServiceImpl) ProviderService {
	return &impl
}
