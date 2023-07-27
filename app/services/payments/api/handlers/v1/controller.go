package v1

import (
	"fmt"

	"github.com/lenguti/ezuzu/business/data/db"
	"github.com/lenguti/ezuzu/foundation/api"
	"github.com/rs/zerolog"
)

// Controller - represents our handler service orchestrator.
type Controller struct {
	db     *db.DB
	config Config
	log    zerolog.Logger
	router *api.Router
}

// NewController - initializes a new controller with all its services.
func NewController(log zerolog.Logger, cfg Config) (*Controller, error) {
	ddb, err := db.New(db.Config{
		User:         cfg.DBUser,
		Password:     cfg.DBPass,
		Name:         cfg.DBName,
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("new controller: unable to initialize new db: %w", err)
	}

	return &Controller{db: ddb,
		config: cfg,
		log:    log,
		router: api.NewRouter(),
	}, nil
}
