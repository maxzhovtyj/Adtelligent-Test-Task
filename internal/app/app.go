package app

import (
	delivery "github.com/maxzhovtyj/Adtelligent-Test-Task/internal/delivery/http"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/db/mysqldb"
	"log"
	"net/http"
)

func Run() {
	dbClient, err := mysqldb.NewClient()
	if err != nil {
		return
	}

	_ = repository.New(dbClient)
	_ = service.New()
	handler := delivery.NewHandler()

	log.Fatal(http.ListenAndServe(":8080", handler.Init()))
}
