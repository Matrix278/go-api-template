package repository

import (
	"context"
	"database/sql"
	"go-api-template/model/commonerrors"
	repositorymodel "go-api-template/repository/model"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type IUser interface {
	Begin() (*sqlx.Tx, error)
	SelectUserByFilter(ctx context.Context, filter repositorymodel.UsersFilter) (*repositorymodel.User, error)
}

type user struct {
	db     *sqlx.DB
	tracer trace.Tracer
}

func NewUser(db *sqlx.DB) IUser {
	return &user{
		db:     db,
		tracer: otel.Tracer("repository/user"),
	}
}

func (repository *user) Begin() (*sqlx.Tx, error) {
	return repository.db.Beginx()
}

func (repository *user) SelectUserByFilter(ctx context.Context, filter repositorymodel.UsersFilter) (*repositorymodel.User, error) {
	ctx, span := repository.tracer.Start(ctx, "SelectUserByFilter")
	defer span.End()

	var user repositorymodel.User
	whereCondition, args := buildUsersWhereCondition(filter)

	query := `SELECT * FROM users`
	if whereCondition != "" {
		query += " WHERE " + whereCondition
	}

	span.SetAttributes(
		attribute.String("db.statement", query),
		attribute.String("db.system", "postgresql"),
	)

	if err := repository.db.GetContext(ctx, &user, query, args...); err != nil {
		span.RecordError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, commonerrors.ErrUserNotFound
		}
		return nil, errors.Wrap(err, "selecting user by filter failed")
	}

	return &user, nil
}

func buildUsersWhereCondition(filter repositorymodel.UsersFilter) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	if filter.ID != nil {
		conditions = append(conditions, "id = $1")
		args = append(args, *filter.ID)
	}

	return strings.Join(conditions, " AND "), args
}
