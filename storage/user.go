package storage

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/grps_1/modles"
)

type Storage struct {
	db *pgx.Conn
}

type StorageI interface {
	CreateUser(ctx context.Context,req modles.User) (string, error)
}

func NewStorage(db *pgx.Conn) StorageI {
	return &Storage{db: db}
}

func (s *Storage) CreateUser(ctx context.Context,req modles.User) (string, error) {
	UserId := uuid.New()

	query := `
		INSERT INTO 
			users (
				user_id, 
				name,
				surname, 
				password
			) VALUES (
			 $1, $2, $3, $4
			)`

	_, err := s.db.Exec(
		ctx, query,
		UserId,
		req.Name,
		req.SurName,
		req.Password,
	)

	if err != nil {

		log.Println("error on CreateUser", err)
		return "",err
	}

	return "succesfully created",nil

}
