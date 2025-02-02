package repository

import (
	"database/sql"
	"go-api-template/model/commonerrors"
	repositorymodel "go-api-template/repository/model"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type IUser interface {
	Begin() (*sqlx.Tx, error)
	SelectUserByID(userID strfmt.UUID4) (*repositorymodel.User, error)
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

func (repository *user) SelectUserByID(userID strfmt.UUID4) (*repositorymodel.User, error) {
	var user repositorymodel.User

	if err := repository.db.Get(&user, `
        SELECT
            *
        FROM
            users
        WHERE
            id = $1
    `, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, commonerrors.ErrUserNotFound
		}

		return nil, errors.Wrap(err, "selecting user by ID failed")
	}

	return &user, nil
}
