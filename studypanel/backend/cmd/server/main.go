package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"studypanel-backend/config"
	"studypanel-backend/internal/handler"
	"studypanel-backend/internal/middleware"
	mongoRepo "studypanel-backend/internal/repository/mongodb"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("falha ao conectar no MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("erro ao desconectar MongoDB: %v", err)
		}
	}()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB não responde: %v", err)
	}
	log.Printf("conectado ao MongoDB: %s", cfg.MongoURI)

	db := client.Database(cfg.DBName)
	repo := mongoRepo.NewStudyItemRepository(db)

	createUC := usecase.NewCreateStudyItemUseCase(repo)
	getUC := usecase.NewGetStudyItemUseCase(repo)
	listUC := usecase.NewListStudyItemsUseCase(repo)
	updateUC := usecase.NewUpdateStudyItemUseCase(repo)
	deleteUC := usecase.NewDeleteStudyItemUseCase(repo)

	h := handler.NewStudyItemHandler(createUC, getUC, listUC, updateUC, deleteUC)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"studypanel-backend"}`))
	})

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      middleware.CORS(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("servidor iniciado na porta %s (env: %s)", cfg.ServerPort, cfg.AppEnv)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}
