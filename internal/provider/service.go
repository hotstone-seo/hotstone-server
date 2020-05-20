package provider

import (
	"context"
	"net/url"

	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/internal/analyt"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

type (
	// Service contain logic for provider api
	// @mock
	Service interface {
		Match(context.Context, url.Values) (*MatchResponse, error)
		FetchTags(context.Context, url.Values) ([]*ITag, error)
		FetchTagsWithCache(context.Context, url.Values, *cachekit.Pragma) ([]*ITag, error)
	}

	// ServiceImpl is implementation of Provider
	ServiceImpl struct {
		dig.In
		analyt.RuleMatchingRepo
		repository.DataSourceRepo
		repository.RuleRepo
		repository.TagRepo
		repository.StructuredDataRepo

		Redis *redis.Client
		urlstore.Store
	}

	// ITag is tag after interpolate with data
	ITag repository.Tag

	// IDataSource is datasource after interpolate with data
	IDataSource repository.DataSource
)

// NewService return new instance of Service
// @constructor
func NewService(impl ServiceImpl) Service {
	return &impl
}