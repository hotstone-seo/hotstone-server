package service

import (
	"context"
	"encoding/json"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"go.uber.org/dig"
)

type (
	// AuditTrailService contain logic for AuditTrail Controller
	// @mock
	AuditTrailService interface {
		Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error)
		RecordChanges(ctx context.Context, record Record) (lastInsertID int64, err error)
	}

	// AuditTrailServiceImpl is implementation of AuditTrailService
	AuditTrailServiceImpl struct {
		dig.In
		AuditTrailRepo repository.AuditTrailRepo
	}

	// Record represents operations that will be logged into audit trail
	Record struct {
		EntityName string
		EntityID   int64
		Operation  OperationType
		PrevData   interface{}
		NextData   interface{}
	}

	// OperationType is type of changes operation
	OperationType string
)

const (
	Insert OperationType = "INSERT"
	Update               = "UPDATE"
	Delete               = "DELETE"
)

// NewAuditTrailService return new instance of AuditTrailService
// @constructor
func NewAuditTrailService(impl AuditTrailServiceImpl) AuditTrailService {
	return &impl
}

// Find audit trail data
func (r *AuditTrailServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.AuditTrail, error) {
	return r.AuditTrailRepo.Find(ctx, paginationParam)
}

// RecordChanges insert changes
func (r *AuditTrailServiceImpl) RecordChanges(ctx context.Context, record Record) (lastInsertID int64, err error) {
	prevDataJSON := repository.JSON("{}")
	if record.PrevData != nil {
		prevDataJSON, err = json.Marshal(record.PrevData)
		if err != nil {
			return
		}
	}

	nextDataJSON := repository.JSON("{}")
	if record.NextData != nil {
		nextDataJSON, err = json.Marshal(record.NextData)
		if err != nil {
			return
		}
	}

	auditTrail := repository.AuditTrail{
		EntityName: record.EntityName,
		EntityID:   record.EntityID,
		Operation:  string(record.Operation),
		Username:   repository.GetUsername(ctx),
		OldData:    prevDataJSON,
		NewData:    nextDataJSON,
	}

	return r.AuditTrailRepo.Insert(ctx, auditTrail)
}
