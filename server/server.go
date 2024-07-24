package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jasurxaydarov/grps_1/config"
	"github.com/jasurxaydarov/grps_1/modles"
	"github.com/jasurxaydarov/grps_1/pgx/db"
	"github.com/jasurxaydarov/grps_1/protos"
	"github.com/jasurxaydarov/grps_1/storage"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
)

type service struct {
	protos.UnimplementedUserServiceServer
}

func (s *service) UserApi(ctx context.Context, in *protos.UserRequest) (*protos.UserResponse, error) {

	cfg := config.Load()
	pgcfg, err := db.ConnToDb(cfg.PgConfig)

	if err != nil {
		logger.Error(err)
		return nil, nil
	}
	fmt.Println(pgcfg)

	str := storage.NewStorage(pgcfg)

	resp, err := str.CreateUser(ctx, modles.User{Name: in.Name, SurName: in.Surname, Password: in.Password})

	if err != nil {

		logger.Error(err)
		return nil, err
	}

	fmt.Println(in)

	return &protos.UserResponse{Resp: resp}, nil
}

func main() {

	listen, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	protos.RegisterUserServiceServer(s, &service{})

	fmt.Println("service running on :8000")

	err = s.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}

}
