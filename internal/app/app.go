package app

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/config"
	delivery "github.com/maxzhovtyj/Adtelligent-Test-Task/internal/delivery/http"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/server"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/db/mysqldb"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/hash"
	"log"
)

func Run() {
	log.Println("initializing config")
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("failed to initialize config, %v", err)
	}

	log.Println("initializing database mySQL client")
	dbClient, err := mysqldb.NewClient(&cfg.DB)
	if err != nil {
		log.Fatalf("failed while connecting to database, %v", err)
	}

	log.Println("initializing token manager")
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		log.Fatalf("failed to initialize token manager, %v", err)
	}

	log.Println("initializing hashing")
	hashing := hash.NewSHA1Hashing(cfg.Auth.PasswordSalt)
	if err != nil {
		log.Fatalf("failed tot initialize hashing manager, %v", err)
	}

	log.Println("initializing repository, service and handler")
	repo := repository.New(dbClient)
	services := service.New(repo, tokenManager, cfg.Auth.JWT.AccessTokenTTL, cfg.Auth.JWT.RefreshTokenTTL, hashing)
	handler := delivery.NewHandler(services, tokenManager)

	srv := server.New(cfg, handler.Init())

	log.Printf("Run application on port %s", cfg.HTTP.Port)
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
