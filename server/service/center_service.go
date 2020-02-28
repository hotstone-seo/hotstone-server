package service

import (
	"context"
	"fmt"
	"time"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbtype"
	"go.uber.org/dig"
)

// CenterService is center related logic [mock]
type CenterService interface {
	AddMetaTag(ctx context.Context, req AddMetaTagRequest) (int64, error)
	AddTitleTag(ctx context.Context, req AddTitleTagRequest) (int64, error)
	AddCanonicalTag(ctx context.Context, req AddCanonicalTagRequest) (int64, error)
	AddScriptTag(ctx context.Context, req AddScriptTagRequest) (int64, error)
}

// CenterServiceImpl implementation of CenterService
type CenterServiceImpl struct {
	dig.In
	repository.TagRepo
}

// NewCenterService return new instance of CenterService [constructor]
func NewCenterService(impl CenterServiceImpl) CenterService {
	return &impl
}

// AddMetaTag to add metaTag
func (i *CenterServiceImpl) AddMetaTag(ctx context.Context, req AddMetaTagRequest) (lastInsertedID int64, err error) {
	lastInsertedID, err = i.TagRepo.Insert(ctx, repository.Tag{
		RuleID:     req.RuleID,
		Locale:     req.Locale,
		Type:       "meta",
		Attributes: dbtype.JSON(fmt.Sprintf(`{"name":"%s", "content":"%s"}`, req.Name, req.Content)),
		Value:      "",
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
	return
}

// AddTitleTag to add titleTag
func (i *CenterServiceImpl) AddTitleTag(ctx context.Context, req AddTitleTagRequest) (lastInsertedID int64, err error) {
	lastInsertedID, err = i.TagRepo.Insert(ctx, repository.Tag{
		RuleID:     req.RuleID,
		Locale:     req.Locale,
		Type:       "title",
		Attributes: dbtype.JSON(`{}`),
		Value:      req.Title,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
	return
}

// AddCanonicalTag to add canonicalTag
func (i *CenterServiceImpl) AddCanonicalTag(ctx context.Context, req AddCanonicalTagRequest) (lastInsertedID int64, err error) {
	lastInsertedID, err = i.TagRepo.Insert(ctx, repository.Tag{
		RuleID:     req.RuleID,
		Locale:     req.Locale,
		Type:       "link",
		Attributes: dbtype.JSON(fmt.Sprintf(`{"href":"%s","rel":"canonical"}`, req.Href)),
		Value:      "",
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
	return
}

// AddScriptTag to add scriptTag
func (i *CenterServiceImpl) AddScriptTag(ctx context.Context, req AddScriptTagRequest) (lastInsertedID int64, err error) {
	lastInsertedID, err = i.TagRepo.Insert(ctx, repository.Tag{
		RuleID:     req.RuleID,
		Locale:     req.Locale,
		Type:       "script",
		Attributes: dbtype.JSON(fmt.Sprintf(`{"source":"%s"}`, req.Source)),
		Value:      req.Type,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
	return
}