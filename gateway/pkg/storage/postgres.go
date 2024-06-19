package storage

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

const (
	_defaultConnectionAttempts = 10
	_defaultConnectionTimeout  = time.Second
)

type Postgres interface {
	Stats(ctx context.Context) *pgxpool.Stat
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

type Pool struct {
	db *pgxpool.Pool
}

func Connect(connectionUrl string) (Postgres, error) {
	connectionAttempts := _defaultConnectionAttempts
	var result *pgxpool.Pool
	var err error
	for connectionAttempts > 0 {
		result, err = pgxpool.Connect(context.Background(), connectionUrl)
		if err == nil {
			break
		}

		log.Printf("ATTEMPT %d TO CONNECT TO POSTGRES BY URL FAILED: %s\n", connectionAttempts, err.Error())

		connectionAttempts--

		time.Sleep(_defaultConnectionTimeout)
	}

	if result == nil {
		log.Printf("POSTGRES CONNECTION ERROR: %v\n", err)
		return nil, err
	}

	return &Pool{db: result}, nil
}

func (p Pool) Stats(_ context.Context) *pgxpool.Stat {
	return p.db.Stat()
}

func (p Pool) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.db.Begin(ctx)
}

func (p Pool) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return p.db.Query(ctx, query, args[:]...)
}

func (p Pool) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := p.db.Query(ctx, query, args[:]...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanOne(dest, rows)
}

func (p Pool) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := p.db.Query(ctx, query, args[:]...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanAll(dest, rows)
}

func (p Pool) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return p.db.Exec(ctx, query, args[:]...)
}

func (p Pool) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return p.db.QueryRow(ctx, query, args[:]...)
}
