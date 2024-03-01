package main

import (
	casestudyapi "github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/case-study-api"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/prepare"
	sqliterepo "github.com/AtakanPehlivanoglu/midas-case-study-api/internal/infra/sqlite"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/usecase/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./assets/file.db"
)

func main() {
	logger := prepare.SLogger()

	// init chi router
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	sqliteDb, err := prepare.SqliteDb(dbDriverName, dbName)
	if err != nil {
		log.Fatalf("error on preparing sqliteDb, %v", err)
	}

	defer sqliteDb.Close()

	sqliteRepository, err := sqliterepo.NewSqliteRepository(sqliterepo.SqliteRepositoryArgs{
		Db: sqliteDb,
	})
	if err != nil {
		log.Fatalf("error on NewSqliteRepository, %v", err)
	}

	// ideally this should be retrieved from app configuration
	shredConfig := prepare.ShredConfig(3, false, false)

	// handlers
	handlerShred, err := handlers.NewShred(handlers.NewShredArgs{
		Logger:      logger,
		ShredConfig: shredConfig,
	})
	if err != nil {
		log.Fatalf("error on NewShred, %v", err)
	}

	handlerFillData, err := handlers.NewFillData(handlers.NewFillDataArgs{
		Logger: logger,
	})
	if err != nil {
		log.Fatalf("error on NewFillData, %v", err)
	}

	handlerDumpDb, err := handlers.NewDumpDb(handlers.NewDumpDbArgs{
		Logger:     logger,
		Repository: sqliteRepository,
	})
	if err != nil {
		log.Fatalf("error on NewDumpDb, %v", err)
	}

	implementation, err := casestudyapi.NewCaseStudyAPI(
		casestudyapi.NewCaseStudyAPIArgs{
			Logger:          logger,
			ShredHandler:    handlerShred,
			FillDataHandler: handlerFillData,
			DumpDbHandler:   handlerDumpDb,
		})

	if err != nil {
		log.Fatalf("error on NewCaseStudyAPI, %v", err)
	}

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get(casestudyapi.HealthEndpoint, casestudyapi.Health)
		r.Route("/file", func(r chi.Router) {
			r.Delete("/shred/{filePath}", implementation.Shred)
			r.Post("/fill", implementation.FillData)
			r.Post("/dump", implementation.DumpDb)
		})
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
