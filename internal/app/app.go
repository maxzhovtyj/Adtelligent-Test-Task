package app

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/config"
	delivery "github.com/maxzhovtyj/Adtelligent-Test-Task/internal/delivery/http"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/db/mysqldb"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/hash"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("failed to initialize config, %v", err)
	}

	dbClient, err := mysqldb.NewClient(cfg.DB.User, cfg.DB.Password, cfg.DB.Database)
	if err != nil {
		return
	}

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		log.Fatalf("failed to initialize token manager, %v", err)
	}

	hashing := hash.NewSHA1Hashing(cfg.Auth.PasswordSalt)
	if err != nil {
		log.Fatalf("failed tot initialize hashing manager, %v", err)
	}

	repo := repository.New(dbClient)
	_ = service.New(repo, tokenManager, cfg.Auth.JWT.AccessTokenTTL, cfg.Auth.JWT.RefreshTokenTTL, hashing)
	handler := delivery.NewHandler()

	log.Fatal(http.ListenAndServe(":8080", handler.Init()))
}
