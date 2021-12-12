package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	grpcAddr    string
	gismeteoURL string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		grpcAddrEnv    = "GRPC_SERVER_ADDR"
		gismeteoURLEnv = "GISMETEO_URL"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.grpcAddr, ok = os.LookupEnv(grpcAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, grpcAddrEnv)
	}

	cfg.gismeteoURL, ok = os.LookupEnv(gismeteoURLEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, gismeteoURLEnv)
	}

	return cfg, nil
}
