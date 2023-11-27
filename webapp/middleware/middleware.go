package middleware

import (
	"context"
	"fmt"
	"monban/database"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type TxHandler func(http.ResponseWriter, *http.Request, pgx.Tx)

func Transaction(db *database.DB, handlerFunc TxHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Failed to open transaction: %s \n", err.Error())
			return
		}

		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tx.Rollback(context.Background())
				return
			} else if r != nil {
				tx.Rollback(context.Background())
				return
			}
			err = tx.Commit(context.Background())
		}()

		handlerFunc(w, r, tx)
	})
}
