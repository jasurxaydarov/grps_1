package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jasurxaydarov/grps_1/modles"
	"github.com/jasurxaydarov/grps_1/protos"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
 var req modles.User
	


	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}


	client := protos.NewUserServiceClient(conn)

	fmt.Println("enter first name ")
	fmt.Scanln(&req.Name)

	fmt.Println("enter secound name ")
	fmt.Scanln(&req.SurName)

	fmt.Println("enter  password ")
	fmt.Scanln(&req.Password)

	resp,err:=client.UserApi(context.Background(), &protos.UserRequest{Name: req.Name,Surname: req.SurName,Password: req.Password})

	if err!= nil{
		logger.Error(err)
		return
	}
	
	log.Println(resp)
}
