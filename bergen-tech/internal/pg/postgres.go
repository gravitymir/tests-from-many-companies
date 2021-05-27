package postgres

import (
	"context"
	"errors"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/config"
	pgx "github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

var PgStatusWorkOk bool

func Init() error {
	if config.Get().PgURL == "" {
		return errors.New("PgURL is empty, set PG_URL in file \".env\"\n")
	}

	conn, err := pgx.Connect(context.Background(), config.Get().PgURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	var result int
	if err := conn.QueryRow(context.Background(), "SELECT 1").Scan(&result); err != nil {
		return err
	} else if result == 1 {
		log.Info("Postgres acces is OK")
		PgStatusWorkOk = true
	}
	return nil
}

func Query() {
	conn, err := pgx.Connect(context.Background(), config.Get().PgURL)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(context.Background(), "select generate_series(1,$1)", 10)

	if err != nil {
		log.Error(err)
	}

	defer rows.Close()

}
