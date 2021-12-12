package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/daniilty/gismeteo-weather-gateway/internal/core"
	"github.com/daniilty/gismeteo-weather-gateway/internal/server"
	schema "github.com/daniilty/weather-gateway-schema"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	service := core.NewServiceImpl(cfg.gismeteoURL)

	loggerCfg := zap.NewProductionConfig()

	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}

	sugaredLogger := logger.Sugar()

	listener, err := net.Listen("tcp", cfg.grpcAddr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	grpcService := server.NewGRPC(service)
	schema.RegisterGismeteoWeatherGatewayServer(grpcServer, grpcService)

	sugaredLogger.Infow("GRPC server is starting.", "addr", listener.Addr())
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			sugaredLogger.Errorw("Server failed to start.", "err", err)
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan

	sugaredLogger.Info("Gracefully stopping GRPC server.")
	grpcServer.GracefulStop()

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
