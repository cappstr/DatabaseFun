package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cappstr/database_fun/config"
	_ "github.com/lib/pq"
)

func NewPostgresStore(conf config.Config) (*sql.DB, error) {
	dsn := conf.DatabaseUrl()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database connection: %w", err)
	}
	return db, nil
}
