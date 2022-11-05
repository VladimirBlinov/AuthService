package main

import (
	"fmt"
	"log"
	"net"

	"github.com/VladimirBlinov/AuthService/internal/app/authserver"
	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/store/inmem"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	store := inmem.New()
	sm := authserver.NewSessionManager(authservice.UnimplementedAuthServiceServer{}, store)

	authservice.RegisterAuthServiceServer(server, sm)

	fmt.Print("starting server ...")
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
