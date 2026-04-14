package main

import (
	"context"
	"log/slog"
	"myavito/internal/handler"
	"myavito/internal/repository"
	"net/http"
	"os"
)

var (
	ctx = context.Background()
)

func main() {
	conn, err := repository.ConnectDB(ctx)
	if err != nil {
		slog.Error("failed connect to db", "fatal", err)
		os.Exit(1)
	}
	rep := repository.Repo{
		Conn: conn,
	}

	userHand := handler.UserHandler{
		Repo: rep,
		Ctx:  ctx,
	}

	http.HandleFunc("POST /user", userHand.CreateUser)

	slog.Info("start http server")

	http.ListenAndServe(":8081", nil)
}
