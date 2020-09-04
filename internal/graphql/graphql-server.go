package gqlserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/vladimir-kopaliani/message-example/internal/graphql/resolver"
	er "github.com/vladimir-kopaliani/message-example/internal/internal_errors"
	"github.com/vladimir-kopaliani/message-example/internal/models"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Server represents graphQL Server
type Server struct {
	service resolver.Service
	logger  Logger
}

// Configuration contains settings for graphQL server
type Configuration struct {
	Logger  Logger
	Service resolver.Service
}

// Logger ...
type Logger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Fatal(msg ...interface{})
}

// NewGraphQLServer create new instance of graphQL server
func NewGraphQLServer(conf *Configuration) Server {
	if conf == nil || conf.Service == nil {
		panic("service is not set for GraphQL server")
	}
	if conf.Logger == nil {
		panic("Logger is not set")
	}

	s := Server{
		service: conf.Service,
		logger:  conf.Logger,
	}

	return s
}

// GetGraphQLHandler returns handler for graphQL queries
func (s Server) GetGraphQLHandler() http.HandlerFunc {
	srv := handler.New(
		resolver.NewExecutableSchema(
			resolver.Config{
				Resolvers: &resolver.Resolver{
					BroadcastChannel: make(chan *models.Message),
					Service:          s.service,
				},
			},
		),
	)

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{})
	srv.Use(apollotracing.Tracer{})
	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(250)) // TODO: put introspection in exceptions

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		var e er.Error
		if errors.As(err, &e) {
			if e.IsPublic() {
				return &gqlerror.Error{
					Message:    e.Error(),
					Path:       graphql.GetFieldContext(ctx).Path(),
					Extensions: map[string]interface{}{"code": e.Code()},
				}
			}
		}

		er := er.ErrInternalSever
		return &gqlerror.Error{
			Message:    er.Error(),
			Path:       graphql.GetFieldContext(ctx).Path(),
			Extensions: map[string]interface{}{"code": er.Code()},
		}
	})

	return srv.ServeHTTP
}
