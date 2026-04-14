package repository

import (
	"context"
	"log/slog"
	"myavito/internal/dto"

	"github.com/jackc/pgx/v5"
)

type Repo struct {
	Conn *pgx.Conn
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *dto.CreateUserRequest) error
}

func (r *Repo) SaveUser(ctx context.Context, user *dto.CreateUserRequest) error {
	sqlQuery := `
	INSERT INTO avitoApp.users (name,phone_number)
	VALUES($1,$2);
	`

	if _, err := r.Conn.Exec(ctx, sqlQuery, user.Name, user.Phone_number); err != nil {
		slog.Error("failed saved user in db", "error", err)
	}

	slog.Info("User saved")
	return nil
}
