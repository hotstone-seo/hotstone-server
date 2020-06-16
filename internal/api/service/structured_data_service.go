package service

import (
	"context"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// StructuredDataService manages instances of Structured Data
// @mock
type StructuredDataService interface {
	FindByRule(ctx context.Context, ruleID int64) ([]*repository.StructuredData, error)
	repository.StructuredDataRepo
}

// StructuredDataServiceImpl is an impolementation of StructuredDataService
type StructuredDataServiceImpl struct {
	dig.In
	repository.StructuredDataRepo
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewStructuredDataService returns nrw instance of StructuredDataService
// @ctor
func NewStructuredDataService(impl StructuredDataServiceImpl) StructuredDataService {
	return &impl
}

// FindByRule returns list of structured data based on rule ID
func (s *StructuredDataServiceImpl) FindByRule(ctx context.Context, ruleID int64) ([]*repository.StructuredData, error) {
	return s.Find(ctx, dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)))
}

func (s *StructuredDataServiceImpl) Insert(ctx context.Context, strData repository.StructuredData) (newID int64, err error) {
	if strData.Data == nil {
		strData.Data = map[string]interface{}{}
	}
	if newID, err = s.StructuredDataRepo.Insert(ctx, strData); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"structured data",
			newID,
			repository.Insert,
			nil,
			strData,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return newID, nil
}

func (s *StructuredDataServiceImpl) Update(ctx context.Context, strData repository.StructuredData) (err error) {
	var prevStrData *repository.StructuredData
	if prevStrData, err = s.StructuredDataRepo.FindOne(ctx, strData.ID); err != nil {
		return
	}
	if strData.Data == nil {
		strData.Data = make(map[string]interface{}, 0)
	}
	if err = s.StructuredDataRepo.Update(ctx, strData); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"structured data",
			strData.ID,
			repository.Update,
			prevStrData,
			strData,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

func (s *StructuredDataServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var strData *repository.StructuredData
	if strData, err = s.StructuredDataRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.StructuredDataRepo.Delete(ctx, id); err != nil {
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			"structured data",
			id,
			strData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"structured data",
			id,
			repository.Delete,
			strData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}
