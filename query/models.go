// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package query

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ApiKeysSchemaApiKey struct {
	ApiKeyID    int32
	ApiKeyValue string
	CreatedAt   pgtype.Timestamptz
}

type SchemaMigration struct {
	Version int64
	Dirty   bool
}