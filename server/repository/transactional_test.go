package repository_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/typical-go/typical-rest-server/pkg/typpostgres"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
)

func TestTransactional(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	trx := repository.Transactional{
		DB: typpostgres.NewDB(db),
	}
	t.Run("WHEN error occurred before commit", func(t *testing.T) {
		ctx := context.Background()
		mock.ExpectBegin()
		mock.ExpectRollback()
		commitFn := trx.CommitMe(&ctx)
		func(ctx context.Context) {
			dbtxn.SetErrCtx(ctx, errors.New("unexpected-error"))
		}(ctx)
		commitFn()
		require.EqualError(t, dbtxn.ErrCtx(ctx), "unexpected-error")
	})
	t.Run("WHEN panic occurred before commit", func(t *testing.T) {
		ctx := context.Background()
		mock.ExpectBegin()
		fn := trx.CommitMe(&ctx)
		func(ctx context.Context) { // service level
			defer fn()
			dbtxn.SetErrCtx(ctx, fmt.Errorf("some-logic-error"))
			func(ctx context.Context) { // repository level
				panic("something-dangerous")
			}(ctx)
		}(ctx)
		require.EqualError(t, dbtxn.ErrCtx(ctx), "something-dangerous")
	})
	t.Run("WHEN begin error", func(t *testing.T) {
		ctx := context.Background()
		mock.ExpectBegin().WillReturnError(errors.New("some-begin-error"))
		require.EqualError(t, trx.CommitMe(&ctx)(), "some-begin-error")
	})
	t.Run("WHEN commit error", func(t *testing.T) {
		ctx := context.Background()
		mock.ExpectBegin()
		mock.ExpectCommit().WillReturnError(errors.New("some-commit-error"))
		require.EqualError(t, trx.CommitMe(&ctx)(), "some-commit-error")
		require.NoError(t, dbtxn.ErrCtx(ctx))
		require.NotNil(t, dbtxn.TxCtx(ctx, nil))
	})
	t.Run("WHEN rolback error", func(t *testing.T) {
		ctx := context.Background()
		mock.ExpectBegin()
		mock.ExpectRollback().WillReturnError(errors.New("some-rollback-error"))
		commitFn := trx.CommitMe(&ctx)
		func(ctx context.Context) {
			dbtxn.SetErrCtx(ctx, errors.New("unexpected-error"))
		}(ctx)
		require.EqualError(t, commitFn(), "some-rollback-error")
		require.EqualError(t, dbtxn.ErrCtx(ctx), "unexpected-error")
	})
}
