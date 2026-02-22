package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"w4ll3t/internal/config"
	"w4ll3t/internal/handler"
	"w4ll3t/internal/repository"
	"w4ll3t/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewWalletRepository(pool)
	svc := service.NewWalletService(repo)
	h := handler.NewWalletHandler(svc)

	r := chi.NewRouter()
	r.Post("/api/v1/wallet", h.UpdateWallet)
	r.Get("/api/v1/wallets/{id}", h.GetBalance)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	go func() {
		log.Println("Server started on port", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = srv.Shutdown(ctxShutdown)
}
