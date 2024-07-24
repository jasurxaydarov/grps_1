package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/grps_1/config"
)

func ConnToDb(pgCfg config.PgConfig) (*pgx.Conn, error) {

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)

	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Println("eror on ConnToDb", err)
		return nil, nil
	}
	return conn, nil
}
