package boot

import (
	"job-order-be/docs"
	"job-order-be/internal/data/auth"
	"job-order-be/pkg/httpclient"
	"job-order-be/pkg/tracing"
	"log"
	"net/http"

	"job-order-be/internal/config"
	jaegerLog "job-order-be/pkg/log"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	joborderData "job-order-be/internal/data/joborder"
	joborderServer "job-order-be/internal/delivery/http"
	joborderHandler "job-order-be/internal/delivery/http/joborder"
	joborderService "job-order-be/internal/service/joborder"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg := config.Get()
	// Open MySQL DB Connection
	db, err := openDatabases(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	//
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	// Set logger used for jaeger
	logger, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "joborder"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("joborder", zlogger)
	defer closer.Close()

	httpc := httpclient.NewClient(tracer)
	ad := auth.New(httpc, cfg.API.Auth)

	// Diganti dengan domain yang anda buat
	sd := joborderData.New(db, tracer, zlogger)
	ss := joborderService.New(sd, ad, tracer, zlogger)
	sh := joborderHandler.New(ss, tracer, zlogger)

	config.PrepareWatchPath()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := config.Init()
		if err != nil {
			log.Printf("[VIPER] Error get config file, %v", err)
		}
		cfg := config.Get()
		masterNew, err := openDatabases(cfg)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize database connection: %v", err)
		} else {
			*db = *masterNew
			sd.InitStmt()
		}

	})
	s := joborderServer.Server{
		Joborder: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func openDatabases(cfg *config.Config) (master *sqlx.DB, err error) {
	master, err = openConnectionPool("mysql", cfg.Database.Master)
	if err != nil {
		return master, err
	}

	return master, err
}

func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(driver, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
