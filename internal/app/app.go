package app

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/config"
	delivery "github.com/maxzhovtyj/Adtelligent-Test-Task/internal/delivery/http"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/db/mysqldb"
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

	_ = repository.New(dbClient)
	_ = service.New()
	handler := delivery.NewHandler()

	log.Fatal(http.ListenAndServe(":8080", handler.Init()))
}
