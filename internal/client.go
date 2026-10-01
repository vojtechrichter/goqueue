package internal

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HandlerFunc func(ctx context.Context, job *Job) error

type Client struct {
	handlers   map[string]HandlerFunc
	dbConnPool *pgxpool.Pool
}

func newDbConnectionPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to create DB connection pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("Failed to ping DB: %w", err)
	}
	return pool, nil
}

func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("LoadConfig: %w", err)
	}

	connPool, err := newDbConnectionPool(ctx, cfg.Dsn())
	if err != nil {
		return nil, fmt.Errorf("newDbConnectionPool: %w", err)
	}

	return &Client{
		handlers:   make(map[string]HandlerFunc),
		dbConnPool: connPool,
	}, nil
}

func (c *Client) Enqueue(ctx context.Context) {
}

func (c *Client) Start(ctx context.Context) {
}

func (c *Client) Stop(ctx context.Context) {
}
