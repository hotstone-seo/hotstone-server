package service

import (
	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"go.uber.org/dig"
)

// MetricsRuleMatchingService contain logic for MetricsUnmatchedController [mock]
type MetricsRuleMatchingService interface {
	repository.MetricsRuleMatchingRepo

	SetMatched(m *repository.MetricsRuleMatching, matchedURL string, ruleID int64)
	SetMismatched(m *repository.MetricsRuleMatching, mismatchedURL string)
}

// MetricsRuleMatchingServiceImpl is implementation of MetricsRuleMatchingService
type MetricsRuleMatchingServiceImpl struct {
	dig.In
	repository.MetricsRuleMatchingRepo
}

// NewMetricsRuleMatchingService return new instance of MetricsRuleMatchingService [constructor]
func NewMetricsRuleMatchingService(impl MetricsRuleMatchingServiceImpl) MetricsRuleMatchingService {
	return &impl
}

// SetMatched to set matched
func (s *MetricsRuleMatchingServiceImpl) SetMatched(m *repository.MetricsRuleMatching, matchedURL string, ruleID int64) {
	m.IsMatched = 1
	m.URL = &matchedURL
	m.RuleID = &ruleID
}

// SetMismatched to set mismatched
func (s *MetricsRuleMatchingServiceImpl) SetMismatched(m *repository.MetricsRuleMatching, mismatchedURL string) {
	m.IsMatched = 0
	m.URL = &mismatchedURL
}
