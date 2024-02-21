package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"projectAuth/sso/internal/app"
	"projectAuth/sso/internal/config"
	"projectAuth/sso/internal/lib/logger/handlers/slogpretty"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "Dev"
	envProd  = "Prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	//log.Info("starting application", slog.Any("cfg", cfg))

	log.Info("starting application")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL) //General

	go application.GRPCSrv.MustRun()

	//TODO: иницилизация приложения (app)

	//TODO: запустить gRPC-сервер приложение

	//Grascefull shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signStop := <-stop //Waiting shutdown of app

	log.Info("stopping application", slog.String("signal", signStop.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)

}
