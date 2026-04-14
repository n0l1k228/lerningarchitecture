package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed connect db: %w", err)
	}
	userDB := os.Getenv("POSTGRES_USER")
	passDB := os.Getenv("POSTGRES_PASSWORD")
	DB := os.Getenv("POSTGRES_DB")
	wayDB := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", userDB, passDB, DB)
	conn, err := pgx.Connect(ctx, wayDB)
	if err != nil {
		return nil, fmt.Errorf("failed connect db: %w", err)
	}
	return conn, nil
}
