package repo

import (
	"monban/query"

	"github.com/jackc/pgx/v5"
)

type IAuthRepo interface{}

type AuthRepo struct {}

func NewAuthRepo() AuthRepo {
    return AuthRepo{}
}

func (a AuthRepo) GetAllAPIKeys(tx pgx.Tx) {
    query.
}
