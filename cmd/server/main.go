package main

import (
	"fmt"
	"log"
	"net"

	"github.com/VladimirBlinov/AuthService/internal/app/authserver"
	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/store/inmem"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := logrus.New()

	listen, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	store := inmem.New()
	sm := authserver.NewSessionManager(authservice.UnimplementedAuthServiceServer{}, store, logger)

	authservice.RegisterAuthServiceServer(server, sm)

	fmt.Print("starting server ...")
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
