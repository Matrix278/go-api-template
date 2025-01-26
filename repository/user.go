package repository

import (
	"database/sql"
	"go-api-template/model"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type IUser interface {
	SelectUserByID(userID strfmt.UUID4) (*model.User, error)
}

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) IUser {
	return &User{
		db: db,
	}
}

func (repository *User) SelectUserByID(userID strfmt.UUID4) (*model.User, error) {
	var user model.User

	if err := repository.db.Get(&user, `
        SELECT
            *
        FROM
            users
        WHERE
            id = $1
    `, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrUserNotFound
		}

		return nil, errors.Wrap(err, "selecting user by ID failed")
	}

	return &user, nil
}
