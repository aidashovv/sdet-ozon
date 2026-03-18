package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sdet-ozon/internal/mock_scenarios/application"
	pgRepo "sdet-ozon/internal/mock_scenarios/infrastructure/postgres"
	"sdet-ozon/internal/mock_scenarios/presentation"
	"sdet-ozon/internal/pkg/http/server"
	pgPool "sdet-ozon/internal/pkg/postgres"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt, syscall.SIGTERM,
	)
	defer cancel()

	dbCfg := pgPool.LoadConfig()
	srvCfg := server.LoadConfig()

	pool, err := pgPool.NewConnectionPool(ctx, dbCfg)
	if err != nil {
		log.Fatalf("failed to init postgres pool: %v", err)
	}
	defer pool.Close()

	repo := pgRepo.NewMockRepository(pool)
	setupService := application.NewSetupService(repo)
	rateService := application.NewRateService(repo)
	h := presentation.NewHandler(setupService, rateService)
	srv := server.NewHTTPServer(
		srvCfg,
		http.HandlerFunc(h.SetupHandler),
		http.HandlerFunc(h.GetCbrXmlHandler),
		http.HandlerFunc(h.DeleteHandler),
	)

	log.Printf("starting http server on %s", srvCfg.Addr)
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server fatal error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down gracefully")

	if err := srv.Stop(context.Background()); err != nil {
		log.Printf("failed to stop server: %v", err)
	}

	log.Println("server stopped")
}
