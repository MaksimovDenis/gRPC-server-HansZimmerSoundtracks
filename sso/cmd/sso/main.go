package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	ssov1 "projectAuth/protos/gen/go/sso"
	"projectAuth/sso/internal/app"
	"projectAuth/sso/internal/config"
	"projectAuth/sso/internal/lib/logger/handlers/slogpretty"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
)

const (
	envLocal       = "local"
	envDev         = "Dev"
	envProd        = "Prod"
	emptyAppId     = 0
	appID          = 1
	appSecret      = "test-secret"
	passDefaultLen = 10
	deltaSeconds   = 1
	start          = "go run sso/cmd/sso/main.go --config=sso/config/local.yaml"
)

// gRPC token
var tokenExpiration time.Time
var token string

// Background images
var imageURLs []string

// Soundtracks
var interstellarSoundrackUrls []string
var batmanSoundrackUrls []string
var duneSoundrackUrls []string
var inceptionSoundrackUrls []string
var piratesOfTheCaribbeanSoundrackUrls []string

type Client struct {
	api ssov1.AuthClient
	log *slog.Logger
}

func New(ctx context.Context,
	log *slog.Logger,
	addr string,
	timeout time.Duration,
	retriesCount int,
) (*Client, error) {
	const op = "grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{
		api: ssov1.NewAuthClient(cc),
	}, nil
}

func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	if r.Method == "POST" {

		// Чтение данных из тела запроса
		var requestData map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		// Получение значений email и password
		email, ok := requestData["email"].(string)
		if !ok {
			http.Error(w, "Email field is missing or invalid", http.StatusBadRequest)
			return
		}
		fmt.Println(email)

		password, ok := requestData["password"].(string)
		if !ok {
			http.Error(w, "Password field is missing or invalid", http.StatusBadRequest)
			return
		}
		fmt.Println(password)

		// Создание экземпляра Client
		client, err := New(context.Background(), logger, "localhost:44044", time.Second, 3)
		if err != nil {
			http.Error(w, "Failed to create gRPC client", http.StatusInternalServerError)
			return
		}

		// Вызов метода для логина с использованием клиента
		ctx := context.Background()
		respReg, err := client.api.Register(ctx, &ssov1.RegisterRequest{
			Email:    email,
			Password: password,
		})
		if err != nil {
			http.Error(w, "Failed to login", http.StatusInternalServerError)
			return
		}
		fmt.Println(respReg)

	}

}

func handleLogin(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	if r.Method == "POST" {

		// Чтение данных из тела запроса
		var requestData map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		// Получение значений email и password
		email, ok := requestData["email"].(string)
		if !ok {
			http.Error(w, "Email field is missing or invalid", http.StatusBadRequest)
			return
		}
		fmt.Println(email)

		password, ok := requestData["password"].(string)
		if !ok {
			http.Error(w, "Password field is missing or invalid", http.StatusBadRequest)
			return
		}
		fmt.Println(password)

		// Создание экземпляра Client
		client, err := New(context.Background(), logger, "localhost:44044", time.Second, 3)
		if err != nil {
			http.Error(w, "Failed to create gRPC client", http.StatusInternalServerError)
			return
		}

		// Вызов метода для логина с использованием клиента
		ctx := context.Background()
		respLogin, err := client.api.Login(ctx, &ssov1.LoginRequest{
			Email:    email,
			Password: password,
			AppId:    appID,
		})
		if err != nil {
			http.Error(w, "Failed to login", http.StatusInternalServerError)
			return
		}

		token = respLogin.GetToken()
		tokenExpiration = time.Now().Add(time.Second * 10)
	}

}

func Dial(s string, dialOption grpc.DialOption) {
	panic("unimplemented")
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("sso/cmd/templates/index.html")
	tmpl.ExecuteTemplate(w, "index", nil)
}

func interstellarHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case token == "":
		http.Error(w, "Unautorized", http.StatusUnauthorized)
	case time.Now().After(tokenExpiration):
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	default:
		imageInterstellar := imageURLs[3]
		tmpl, _ := template.ParseFiles("sso/cmd/templates/header.html", "sso/cmd/templates/interstellar.html", "sso/cmd/templates/player.html")
		tmpl.ExecuteTemplate(w, "interstellar", imageInterstellar)
	}
}

func trackInterstellarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(interstellarSoundrackUrls)
}

func handlRequest(log *slog.Logger) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("sso/cmd/static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handleRegister(w, r, log) // передача логгера в функцию handleLogin
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handleLogin(w, r, log)
	})

	// Добавьте tokenMiddleware перед вашим обработчиком batmanHandler
	http.HandleFunc("/interstellar", interstellarHandler)
	http.HandleFunc("/interstellarSountrack", trackInterstellarHandler)

	log.Info("starting web-server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Failed to start HTTP server", slog.String("error", err.Error()))
	}
}

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg))

	log.Info("starting application")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL) //General

	go application.GRPCSrv.MustRun()

	//Yandex storage

	//Start HTTP server to serve HTML page
	handlRequest(log)

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
