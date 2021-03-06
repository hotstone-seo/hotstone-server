package service

import (
	"context"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

type (
	// TagService provides available method to be used for managing Tag entity
	// @mock
	TagService interface {
		FindByRuleAndLocale(ctx context.Context, ruleID int64, locale string) ([]*repository.Tag, error)
		repository.TagRepo
	}
	// TagServiceImpl is the implementation of TagService
	TagServiceImpl struct {
		dig.In
		repository.TagRepo
		AuditTrail AuditTrailSvc
	}
)

// NewTagService return new instance of TagService
// @ctor
func NewTagService(impl TagServiceImpl) TagService {
	return &impl
}

// FindByRuleAndLocale returns list of Tag entity by searching on matching ruleID and locale
func (s *TagServiceImpl) FindByRuleAndLocale(ctx context.Context, ruleID int64, locale string) (list []*repository.Tag, err error) {
	return s.Find(ctx,
		dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)),
		dbkit.Equal("locale", locale),
	)
}

// Insert creates a new Tag entity
func (s *TagServiceImpl) Insert(ctx context.Context, tag repository.Tag) (int64, error) {

	if tag.Attributes == nil {
		tag.Attributes = map[string]string{}
	}

	lastInsertedID, err := s.TagRepo.Insert(ctx, tag)
	if err != nil {
		return -1, err
	}

	tag.ID = lastInsertedID

	s.AuditTrail.RecordInsert(ctx, "tags", lastInsertedID, tag)
	return lastInsertedID, nil
}

// Update modify existing Tag entity
func (s *TagServiceImpl) Update(ctx context.Context, tag repository.Tag) (err error) {
	var currentTag *repository.Tag
	if currentTag, err = s.TagRepo.FindOne(ctx, tag.ID); err != nil {
		return
	}

	s.AuditTrail.RecordUpdate(ctx, "tags", tag.ID, currentTag, tag)
	return s.TagRepo.Update(ctx, tag)
}

// Delete tag
func (s *TagServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var currentTag *repository.Tag
	if currentTag, err = s.TagRepo.FindOne(ctx, id); err != nil {
		return
	}
	s.AuditTrail.RecordDelete(ctx, "tags", id, currentTag)
	return s.TagRepo.Delete(ctx, id)
}
