package repo

import (
	"context"
	"fmt"
	apikeyModel "monban/internal/apikey/model"
	"monban/query"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type IAuthRepo interface{}

type AuthRepo struct{}

func NewAuthRepo() AuthRepo {
	return AuthRepo{}
}

func (a AuthRepo) GetAllAPIKeys(tx pgx.Tx) ([]apikeyModel.APIKey, error) {
	q := query.New(tx)
	sqlApiKeys, err := q.GetAPIKeys(context.Background())
	if err != nil {
		return []apikeyModel.APIKey{}, fmt.Errorf("Unable to get all api keys: %v", err)
	}

	apikeys := []apikeyModel.APIKey{}
	for _, sak := range sqlApiKeys {

		apikeys = append(apikeys, apikeyModel.APIKey{
			ID:        int(sak.ApiKeyID),
			Key:       sak.ApiKeyValue,
			CreatedAt: sak.CreatedAt.Time,
		})

	}
	return apikeys, nil
}

func (a AuthRepo) CraeteAPIKey(tx pgx.Tx, ak apikeyModel.APIKey) (int, error) {
	q := query.New(tx)
	id, err := q.CreateAPIKey(context.Background(), query.CreateAPIKeyParams{
		ApiKeyID:    int32(ak.ID),
		ApiKeyValue: ak.Key,
		CreatedAt:   pgtype.Timestamptz{Time: ak.CreatedAt},
	})
	if err != nil {
		return 0, fmt.Errorf("Unable to insert api key: %v", err)
	}

	return int(id), nil
}
