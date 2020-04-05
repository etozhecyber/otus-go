package psql

import (
	"database/sql"

	"github.com/etozhecyber/otus-go/calsrv/utilities"
	_ "github.com/jackc/pgx/stdlib" //psql
)

//PostgresStorage struct
type PostgresStorage struct {
	client *sql.DB
}

//NewPostgresStorage create new psql connection
func NewPostgresStorage(config utilities.Config) (*PostgresStorage, error) {
	db, err := sql.Open("pgx", config.DBDSN)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{
		client: db,
	}, err
}
