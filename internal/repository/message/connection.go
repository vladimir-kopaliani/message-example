package messagerepo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository represents message's repository
type Repository struct {
	isDebugMode bool
	db          *pgxpool.Pool
	logger      Logger
}

// Configuration contains setting for connection to message's repository
type Configuration struct {
	IsDebugMode bool
	Logger      Logger
	Host        string
	Port        string
	Password    string
	DBName      string
	User        string
	SSLMode     string
}

// NewMessageRepository creates new instance of message's repository
func NewMessageRepository(ctx context.Context, conf Configuration) (*Repository, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%s password=%s dbname=%s user=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.Password,
		conf.DBName,
		conf.User,
		conf.SSLMode,
	))
	if err != nil {
		conf.Logger.Error(err)
		return nil, err
	}

	config.MaxConns = 10
	config.MaxConnLifetime = 5 * time.Minute
	config.HealthCheckPeriod = 2 * time.Minute

	if conf.IsDebugMode {
		config.ConnConfig.Logger = logger{conf.Logger}
		config.ConnConfig.LogLevel = pgx.LogLevelDebug
	}

	db, err := pgxpool.ConnectConfig(
		ctx,
		config,
	)
	if err != nil {
		conf.Logger.Error(err)
		return nil, err
	}

	conf.Logger.Info("Conencted to DB: " + conf.Host)

	r := Repository{
		isDebugMode: conf.IsDebugMode,
		logger:      conf.Logger,
		db:          db,
	}

	err = r.prepareDatabase(ctx)
	if err != nil {
		conf.Logger.Error(err)
	}

	return &r, nil
}

// Close close connection to message's repository
func (r *Repository) Close(ctx context.Context) error {
	r.logger.Info("Message Repository is closing...")
	defer r.logger.Info("Message Repository closed")

	r.db.Close()
	return nil
}

func (r *Repository) prepareDatabase(ctx context.Context) error {
	r.db.Exec(ctx, `CREATE TABLE IF NOT EXISTS messages (
		id					VARCHAR(64)		PRIMARY KEY,
		text				VARCHAR(128)	NOT NULL,
		created_at	TIMESTAMPTZ		NOT NULL
		);`)
	return nil
}
