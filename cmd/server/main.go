package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	gqlserver "github.com/vladimir-kopaliani/message-example/internal/graphql"
	httpserver "github.com/vladimir-kopaliani/message-example/internal/http_server"
	"github.com/vladimir-kopaliani/message-example/internal/logger"
	messagerepo "github.com/vladimir-kopaliani/message-example/internal/repository/message"
	"github.com/vladimir-kopaliani/message-example/internal/service"

	"github.com/99designs/gqlgen/graphql/playground"
)

var (
	isDebugMode = true
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// handle interrupt signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh:
			log.Println("Got Interrupt signal. Stutting down...")
			cancel()
		}
	}()

	// set production mode
	if _, ok := os.LookupEnv("PRODUCTION_MODE"); ok {
		isDebugMode = false
	}

	// Logger
	l, err := logger.NewLogger(ctx, isDebugMode)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	// Repository
	messageRepo, err := messagerepo.NewMessageRepository(ctx,
		messagerepo.Configuration{
			IsDebugMode: isDebugMode,
			Logger:      l,
			User:        getEnv(l, "POSTGRES_USER", "postgres"),
			Password:    getEnv(l, "POSTGRES_PASSWORD", "postgres"),
			DBName:      getEnv(l, "POSTGRES_DB_NAME", "postgres"),
			Host:        getEnv(l, "POSTGRES_HOST", "localhost"),
			Port:        getEnv(l, "POSTGRES_PORT", "5432"),
			SSLMode:     getEnv(l, "POSTGRES_SSL_MODE", "disable"),
		})
	if err != nil {
		return
	}
	defer messageRepo.Close(ctx)

	// Service
	serv := service.NewService(&service.Configuration{
		Logger:             l,
		MessagesRepository: messageRepo,
	})

	// GraphQL
	graphQLServer := gqlserver.NewGraphQLServer(&gqlserver.Configuration{
		Logger:  l,
		Service: serv,
	})

	// HTTP server
	httpServer, err := httpserver.NewHTTPServer(ctx,
		&httpserver.Configuration{
			Logger:  l,
			Address: getEnv(l, "HTTP_ADDRESS", ":3000"),
			Handlers: []httpserver.Handler{
				// graphQL
				{
					Path: "/graphql",
					Handler: httpserver.RetrieveMetadata(
						graphQLServer.GetGraphQLHandler(),
					),
				},
				// graphQL
				{
					Path:    "/playground",
					Handler: playground.Handler("server", "/graphql"),
				},
			},
		})
	if err != nil {
		l.Error(err)
		return
	}

	go httpServer.Launch(ctx)
	defer httpServer.Close(ctx)

	if isDebugMode {
		l.Info("Started in debug mode")
	}

	<-ctx.Done()
}

func getEnv(logger *logger.Logger, env string, defaultValue string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		logger.Infof("For %s applied default value: %q", env, defaultValue)
		return defaultValue
	}
	return value
}
