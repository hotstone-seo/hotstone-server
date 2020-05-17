package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/dchest/uniuri"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// ClientKeyService contain logic for ClientKeyController
// @mock
type ClientKeyService interface {
	repository.ClientKeyRepo
}

// ClientKeyServiceImpl is implementation of ClientKeyService
type ClientKeyServiceImpl struct {
	dig.In
	repository.ClientKeyRepo
	dbtxn.Transactional
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewClientKeyService return new instance of ClientKeyService
// @constructor
func NewClientKeyService(impl ClientKeyServiceImpl) ClientKeyService {
	return &impl
}

// Insert client key
func (s *ClientKeyServiceImpl) Insert(ctx context.Context, data repository.ClientKey) (newData repository.ClientKey, err error) {
	newPrefix, newKey, newKeyHashed := generateClientKey()
	data.Prefix = newPrefix
	data.Key = newKeyHashed
	newData, err = s.ClientKeyRepo.Insert(ctx, data)
	newData.Key = newKey
	if err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"client_keys",
			newData.ID,
			repository.Insert,
			nil,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return
}

// Update client key
func (s *ClientKeyServiceImpl) Update(ctx context.Context, data repository.ClientKey) (err error) {
	var oldData *repository.ClientKey
	if oldData, err = s.ClientKeyRepo.FindOne(ctx, data.ID); err != nil {
		return
	}
	if err = s.ClientKeyRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"client_keys",
			data.ID,
			repository.Update,
			oldData,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// Delete client key
func (s *ClientKeyServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.ClientKey
	if oldData, err = s.ClientKeyRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.ClientKeyRepo.Delete(ctx, id); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			"client_keys",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"client_keys",
			id,
			repository.Delete,
			oldData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

func generateClientKey() (prefix, key, keyHashed string) {
	newPrefix := uniuri.NewLen(7)
	newKey := uniuri.NewLen(32)
	return newPrefix, newKey, fmt.Sprintf("%x", sha256.Sum256([]byte(newKey)))
}