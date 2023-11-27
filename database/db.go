package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Conn *pgxpool.Pool
}

func NewDB(user, password, host, name string, port, poolSize int) DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s pool_max_conns=%d",
		user, password, host, port, name, poolSize,
	)

	fmt.Println(dsn)
	connPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic("Unable to connect to db")
	}
	return DB{
		Conn: connPool,
	}
}

func (d *DB) Close() {
	d.Conn.Close()
}

func (d *DB) Begin() (pgx.Tx, error) {
	return d.Conn.Begin(context.Background())

}
