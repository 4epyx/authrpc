package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repositories"
	"github.com/4epyx/authrpc/services"
	"github.com/4epyx/authrpc/utils"
	"google.golang.org/grpc"
)

// TODO: to clean code
func main() {
	host, ok := os.LookupEnv("SERVER_HOST")
	if !ok {
		host = "localhost"
	}

	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "8080"
	}

	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		panic("not found DB_URL in environment variable")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	db, err := utils.ConnectToDB(context.TODO(), dbUrl)
	if err != nil {
		panic(err)
	}
	err = utils.MigrateTable(context.TODO(), db)
	if err != nil {
		panic(err)
	}

	repo := repositories.NewPgxUserRepository(db)

	pb.RegisterRegisterServiceServer(grpcServer, services.NewRegisterService(repo))
	pb.RegisterLoginServiceServer(grpcServer, services.NewLoginService(repo))
	pb.RegisterUserDataServiceServer(grpcServer, services.NewUserDataService(repo))
	pb.RegisterAuthorizationServiceServer(grpcServer, services.NewAuthorizationService())

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
