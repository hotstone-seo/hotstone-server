package service

import (
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController
type ProviderService interface {
	MatchRule(MatchRuleRequest) (*repository.Rule, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
}

// MatchRuleRequest is request for match rule
type MatchRuleRequest struct {
	Path string `json:"path"`
}

// NewProviderService return new instance of ProviderService
func NewProviderService() ProviderService {
	return &ProviderServiceImpl{}
}

// MatchRule to match rule
func (*ProviderServiceImpl) MatchRule(req MatchRuleRequest) (rule *repository.Rule, err error) {
	dataSourceID := int64(88888)
	rule = &repository.Rule{
		ID:           999999,
		Name:         "sample-rule",
		UrlPattern:   "some-url-pattern",
		DataSourceID: &dataSourceID,
	}
	return
}