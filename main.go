package main

import (
	"io/ioutil"
	"net"
	"os"

	"github.com/chaitanyapantheor/golang-grpc-http/gateway"
	"github.com/chaitanyapantheor/golang-grpc-http/insecure"
	buildsv1 "github.com/chaitanyapantheor/golang-grpc-http/proto/builds/v1"
	usersv1 "github.com/chaitanyapantheor/golang-grpc-http/proto/users/v1"
	"github.com/chaitanyapantheor/golang-grpc-http/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	usersv1.RegisterUserServiceServer(s, server.New())
	buildsv1.RegisterBuildServiceServer(s, server.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
