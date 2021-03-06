package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/dbtxn"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// ClientKey represented  an client_keys entity
type ClientKey struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Prefix    string    `json:"prefix"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ClientKeyRepo to handle client_keys entity
// @mock
type ClientKeyRepo interface {
	FindOne(context.Context, ...dbkit.SelectOption) (*ClientKey, error)
	Find(context.Context, ...dbkit.SelectOption) ([]*ClientKey, error)
	Insert(context.Context, ClientKey) (ClientKey, error)
	Delete(context.Context, int64) error
	Update(context.Context, ClientKey) error
}

// ClientKeyRepoImpl is implementation client key repository
type ClientKeyRepoImpl struct {
	dig.In
	*sql.DB
}

// NewClientKeyRepo return new instance of ClientKeyRepo
// @ctor
func NewClientKeyRepo(impl ClientKeyRepoImpl) ClientKeyRepo {
	return &impl
}

// FindOne clientKey
func (r *ClientKeyRepoImpl) FindOne(ctx context.Context, opts ...dbkit.SelectOption) (e *ClientKey, err error) {
	builder := sq.
		Select(
			"id",
			"name",
			"prefix",
			"key",
			"created_at",
			"updated_at",
		).
		From("client_keys").
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	for _, opt := range opts {
		if builder, err = opt.CompileSelect(builder); err != nil {
			return
		}
	}

	e = new(ClientKey)
	if err = builder.QueryRowContext(ctx).Scan(
		&e.ID,
		&e.Name,
		&e.Prefix,
		&e.Key,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return
}

// Find client_keys
func (r *ClientKeyRepoImpl) Find(ctx context.Context, opts ...dbkit.SelectOption) (list []*ClientKey, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"id",
			"name",
			"prefix",
			"key",
			"created_at",
			"updated_at",
		).
		From("client_keys").
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	for _, opt := range opts {
		if builder, err = opt.CompileSelect(builder); err != nil {
			return
		}
	}

	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()

	list = make([]*ClientKey, 0)
	for rows.Next() {
		var e ClientKey
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Prefix,
			&e.Key,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			return
		}
		list = append(list, &e)
	}
	return
}

// Insert clientKey
func (r *ClientKeyRepoImpl) Insert(ctx context.Context, e ClientKey) (newClientKey ClientKey, err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return
	}

	builder := sq.
		Insert("client_keys").
		Columns(
			"name",
			"prefix",
			"key",
		).
		Values(e.Name, e.Prefix, e.Key).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	err = builder.QueryRowContext(ctx).Scan(
		&newClientKey.ID,
		&newClientKey.Name,
		&newClientKey.Prefix,
		&newClientKey.Key,
		&newClientKey.CreatedAt,
		&newClientKey.UpdatedAt,
	)

	if err != nil {
		txn.SetError(err)
		return
	}
	return
}

// Delete clientKey
func (r *ClientKeyRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return
	}

	builder := sq.
		Delete("client_keys").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if _, err = builder.ExecContext(ctx); err != nil {
		txn.SetError(err)
	}
	return
}

// Update tag
func (r *ClientKeyRepoImpl) Update(ctx context.Context, e ClientKey) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return
	}

	builder := sq.
		Update("client_keys").
		Set("name", e.Name).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if _, err = builder.ExecContext(ctx); err != nil {
		txn.SetError(err)
		return
	}
	return
}

// Validate clientKey
func (clientKey ClientKey) Validate() error {
	validate := validator.New()
	return validate.Struct(clientKey)
}
