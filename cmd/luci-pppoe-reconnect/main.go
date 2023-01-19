package main

import (
	"fmt"
	"os"

	"github.com/117503445/luci-pppoe-reconnect/internal/cfg"
	"github.com/117503445/luci-pppoe-reconnect/internal/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"

	"github.com/imroc/req/v3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := log.NewAtLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer func() {
		err = logger.Sync()
	}()

	cfg.InitConfig()

	logger.Info("Hello world!", zap.String("connector.name", viper.GetString("connector.name")))

	client := req.C().DevMode()
	_, err = client.R().Post("http://httpbin.org/get")
	// fmt.Println(resp)

	return err
}
