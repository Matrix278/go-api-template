package repository

import (
	"database/sql"
	"go-api-template/model/commonerrors"
	repositorymodel "go-api-template/repository/model"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type IUser interface {
	Begin() (*sqlx.Tx, error)
	SelectUserByFilter(filter repositorymodel.UsersFilter) (*repositorymodel.User, error)
}

type user struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) IUser {
	return &user{
		db: db,
	}
}

func (repository *user) Begin() (*sqlx.Tx, error) {
	return repository.db.Beginx()
}

func (repository *user) SelectUserByFilter(filter repositorymodel.UsersFilter) (*repositorymodel.User, error) {
	var user repositorymodel.User

	whereCondition, args := buildUsersWhereCondition(filter)

	query := `
        SELECT
            *
        FROM
            users
    `
	if whereCondition != "" {
		query += " WHERE " + whereCondition
	}

	if err := repository.db.Get(&user, query, args...); err != nil {
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
