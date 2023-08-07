package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repository"
	"github.com/4epyx/authrpc/service"
	"github.com/4epyx/authrpc/util"
	"google.golang.org/grpc"
)

func main() {
	host, port := getHostAndPort()

	if err := setupJwtSecret(); err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	repo, err := getPgxRepo()
	if err != nil {
		panic(err)
	}

	registerServices(grpcServer, repo)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func getHostAndPort() (string /*host*/, string /*port*/) {
	host, ok := os.LookupEnv("SERVER_HOST")
	if !ok {
		host = "localhost"
	}

	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "8080"
	}
	return host, port
}

func setupJwtSecret() error {
	val, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return errors.New("JWT_SECRET not found in environment variables")
	}

	util.JwtSecret = []byte(val)
	return nil
}

func getPgxRepo() (*repository.PgxUserRepository, error) {
	dbUrl, err := getDbUrl()
	if err != nil {
		return nil, err
	}

	db, err := util.ConnectToDB(context.TODO(), dbUrl)
	if err != nil {
		return nil, err
	}

	if err := util.MigrateTable(context.TODO(), db); err != nil {
		return nil, err
	}

	return repository.NewPgxUserRepository(db), nil
}

func getDbUrl() (string, error) {
	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		return "", errors.New("DB_URL not found in environment variables")
	}

	return dbUrl, nil
}

func registerServices(grpcServer *grpc.Server, repo repository.UserRepository) {
	pb.RegisterRegisterServiceServer(grpcServer, service.NewRegisterService(repo))
	pb.RegisterLoginServiceServer(grpcServer, service.NewLoginService(repo))
	pb.RegisterUserDataServiceServer(grpcServer, service.NewUserDataService(repo))
	pb.RegisterAuthorizationServiceServer(grpcServer, service.NewAuthorizationService())

}
