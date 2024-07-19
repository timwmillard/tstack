package main

import (
	"app/handler"
	"app/internal/server"
	"app/internal/version"
	"app/model"
	"app/static"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
)

func main() {
	showVersion := flag.Bool("version", false, "get the server version")
	log := flag.String("log", "text", "log output format - either json or text")
	autoMigrate := flag.Bool("migrate", false, "run the database migration on starup")
	flag.Parse()

	appVersion := version.Get()
	if *showVersion {
		fmt.Println("app ", appVersion)
		os.Exit(0)
	}

	logger := slog.Default()
	if *log == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	slog.Info("App: starting", "version", appVersion)

	// Database setup
	database := "app_db"
	if flag.NArg() > 0 {
		database = flag.Arg(0)
	}
	connString := fmt.Sprintf("database=%s", database)

	pool, err := DBPoolConnect(ctx, connString, *autoMigrate)
	if err != nil {
		slog.Error("Database: pool connection", "error", err)
		os.Exit(1)
	}
	defer DBPooleClose(pool)

	// Load configuration
	config := handler.LoadConfig()
	handler.LogConfig(config)

	// Service setup
	service := model.Service{
		DB:     pool,
		Host:   config.Host,
		Log:    *logger,
		AppEnv: config.AppEnv,
	}

	err = service.RiverStart(ctx)
	if err != nil {
		slog.Error("App: river starting", "error", err)
		os.Exit(1)
	}
	defer func() {
		err = service.RiverStop(ctx)
		if err != nil {
			slog.Error("App: river stopping", "error", err)
		}
		slog.Info("App: river stopped")
	}()

	// Routing setup
	router := mux.NewRouter().StrictSlash(true)

	// Session manager
	sess := scs.New()
	sess.Lifetime = 24 * time.Hour
	sess.Cookie.Secure = config.CookieSecure
	sess.Store = pgxstore.New(pool)

	// Routes
	handler.AppRoutes("/", router, config, service, sess)
	handler.AdminRoutes("/admin", router, config, service, sess)

	// Static server
	static.Server("/static", router)

	// Health Check
	router.HandleFunc("/health", Health)

	// Logging
	// routerWithLogging := handlers.LoggingHandler(os.Stdout, router)

	// Server setup
	// server := NewServer(config.Port, routerWithLogging)
	server := server.New(config.Port, router)
	server.Start()
	defer server.Stop()

	// Wait for CTRL-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// We block here until a CTRL-C / SigInt is received
	// Once received, we exit and the server is cleaned up
	<-sigChan
	slog.Info("App: shutdown request via CTRL-C")

	// DBPooleClose(pool)
	// cancel()
}

func Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
