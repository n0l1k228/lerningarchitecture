package main

import (
	"context"
	"log/slog"
	product_handler "myavito/internal/handler/product"
	user_handler "myavito/internal/handler/user"
	"myavito/internal/logger"
	"myavito/internal/repository"
	"myavito/internal/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

var (
	ctx = context.Background()
)

func main() {
	var logger = logger.Logger{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	conn, err := repository.ConnectDB(ctx)
	if err != nil {
		slog.Error("failed connect to db", "fatal", err)
		os.Exit(1)
	}

	repo := repository.NewPostgresRepo(conn)
	srv := service.NewServiceUser(repo)
	srvProd := service.NewProductService(repo)
	hand := user_handler.NewUserHandler(srv, validator.New(), &logger)
	prodHand := product_handler.NewProductHandler(srvProd, validator.New(), &logger)

	http.HandleFunc("POST /user", hand.CreateUser)
	http.HandleFunc("POST /product", prodHand.CreateProduct)

	slog.Info("start http server")
	http.ListenAndServe(":8081", nil)
}
