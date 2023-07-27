package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rajiv-k/cricket/pkg/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

const (
	logLevel   = "debug"
	listenAddr = "/tmp/cricket.sock"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})

	logrus.Info("starting cricket")

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatalf("invalid loglevel : %v", logLevel)
	}
	logrus.SetLevel(level)

	listener, err := net.Listen("unix", listenAddr)
	if err != nil {
		logrus.Fatalf("could not start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	cricketRuntimeServer := server.NewRuntimeServer()
	cricketImageServer := server.NewImageServer()
	v1.RegisterRuntimeServiceServer(grpcServer, cricketRuntimeServer)
	v1.RegisterImageServiceServer(grpcServer, cricketImageServer)
	reflection.Register(grpcServer)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		grpcServer.GracefulStop()
	}(quitCh)

	logrus.Infof("listening on %v://%v", listener.Addr().Network(), listener.Addr().String())

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Errorf("could not run gRPC server: %v", err)
	}
}
