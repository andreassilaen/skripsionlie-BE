package boot

import (
	"log"
	"net/http"
	"skripsi-online-BE/docs"
	"skripsi-online-BE/internal/data/auth"
	"skripsi-online-BE/pkg/httpclient"
	"skripsi-online-BE/pkg/tracing"

	"skripsi-online-BE/internal/config"
	jaegerLog "skripsi-online-BE/pkg/log"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	skripsionlineData "skripsi-online-BE/internal/data/skirpsionline-BE"
	skripsionlineServer "skripsi-online-BE/internal/delivery/http"
	skripsionlineHandler "skripsi-online-BE/internal/delivery/http/skirpsionline-BE"
	skripsionlineService "skripsi-online-BE/internal/service/skirpsionline-BE"
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
	zapLogger := logger.With(zap.String("service", "skirpsionline-BE"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("skirpsionline-BE", zlogger)
	defer closer.Close()

	httpc := httpclient.NewClient(tracer)
	ad := auth.New(httpc, cfg.API.Auth)

	// Diganti dengan domain yang anda buat
	sd := skripsionlineData.New(db, tracer, zlogger)
	ss := skripsionlineService.New(sd, ad, tracer, zlogger)
	sh := skripsionlineHandler.New(ss, tracer, zlogger)

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
	s := skripsionlineServer.Server{
		SkripsionlineBE: sh,
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
