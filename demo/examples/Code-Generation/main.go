package main

import (
	"context"
	"os"
	"time"

	"github.com/anz-bank/sysl/demo/examples/Code-Generation/implementation"
	simple "github.com/anz-bank/sysl/demo/examples/Code-Generation/simple"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.service.anz/sysl/server-lib/common"
	"github.service.anz/sysl/server-lib/config"
	"github.service.anz/sysl/server-lib/core"
	"github.service.anz/sysl/server-lib/logging"
)

func main() {
	// logconfig
	logConfig := config.LogConfig{
		Level:        logrus.InfoLevel,
		Format:       "color",
		ReportCaller: false,
	}

	// library config
	libraryConfig := config.LibraryConfig{
		Log:       logConfig,
		Profiling: true,
	}

	// logging setup
	logger, err := logging.Logger(os.Stdout, &logConfig)
	if err != nil {
		logger.Fatalf("Invalid logger: %s\n", err.Error())
	}

	ctx := common.LoggerToContext(context.Background(), logger, nil)

	// running server on localhost:3000
	serverConfig := config.CommonHTTPServerConfig{
		Common: config.CommonServerConfig{
			HostName: "localhost",
			Port:     3000,
			TLS:      nil,
		},
		BasePath:     "/",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	// running admin server on localhost:3001
	adminServerConfig := config.CommonHTTPServerConfig{
		Common: config.CommonServerConfig{
			HostName: "localhost",
			Port:     3001,
			TLS:      nil,
		},
		BasePath:     "/",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	// service configuration
	simpleCallback := implementation.Callback{}
	simpleServiceHandler := simple.NewServiceHandler(simpleCallback, &simple.ServiceInterface{})
	simpleServiceRouter := simple.NewServiceRouter(simpleCallback, simpleServiceHandler)

	// rest handler setup
	restHandler := implementation.Handler{}
	restHandler.Router = append(restHandler.Router, simpleServiceRouter)
	restHandler.MAdminServerConfig = &adminServerConfig
	restHandler.ServerConfig = &serverConfig
	restHandler.MLibraryConfig = &libraryConfig

	// prometheus monitoring
	prometheusRegistry := prometheus.NewRegistry()

	// server configuration setup and run here
	err = core.Server(ctx, "example", restHandler, implementation.GRPCHandler{}, logger, prometheusRegistry, nil)
	if err != nil {
		logger.Fatalln(err)
	}

}
