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

	// StudyItem
	repo := mongoRepo.NewStudyItemRepository(db)
	createUC := usecase.NewCreateStudyItemUseCase(repo)
	getUC := usecase.NewGetStudyItemUseCase(repo)
	listUC := usecase.NewListStudyItemsUseCase(repo)
	updateUC := usecase.NewUpdateStudyItemUseCase(repo)
	deleteUC := usecase.NewDeleteStudyItemUseCase(repo)
	h := handler.NewStudyItemHandler(createUC, getUC, listUC, updateUC, deleteUC)

	// StudyNote
	noteRepo := mongoRepo.NewStudyNoteRepository(db)
	noteCreateUC := usecase.NewCreateStudyNoteUseCase(noteRepo)
	noteListUC := usecase.NewListStudyNotesUseCase(noteRepo)
	noteUpdateUC := usecase.NewUpdateStudyNoteUseCase(noteRepo)
	noteDeleteUC := usecase.NewDeleteStudyNoteUseCase(noteRepo)
	noteHandler := handler.NewStudyNoteHandler(noteCreateUC, noteListUC, noteUpdateUC, noteDeleteUC)

	// StudyResource
	resourceRepo := mongoRepo.NewStudyResourceRepository(db)
	resourceCreateUC := usecase.NewCreateStudyResourceUseCase(resourceRepo)
	resourceListUC := usecase.NewListStudyResourcesUseCase(resourceRepo)
	resourceUpdateUC := usecase.NewUpdateStudyResourceUseCase(resourceRepo)
	resourceDeleteUC := usecase.NewDeleteStudyResourceUseCase(resourceRepo)
	resourceHandler := handler.NewStudyResourceHandler(resourceCreateUC, resourceListUC, resourceUpdateUC, resourceDeleteUC)

	// QuizQuestion
	quizRepo := mongoRepo.NewQuizQuestionRepository(db)
	quizCreateUC := usecase.NewCreateQuizQuestionUseCase(quizRepo)
	quizListUC := usecase.NewListQuizQuestionsUseCase(quizRepo)
	quizUpdateUC := usecase.NewUpdateQuizQuestionUseCase(quizRepo)
	quizDeleteUC := usecase.NewDeleteQuizQuestionUseCase(quizRepo)
	quizHandler := handler.NewQuizQuestionHandler(quizCreateUC, quizListUC, quizUpdateUC, quizDeleteUC)

	// StudySession
	sessionRepo := mongoRepo.NewStudySessionRepository(db)
	sessionCreateUC := usecase.NewCreateStudySessionUseCase(sessionRepo)
	sessionListUC := usecase.NewListStudySessionsUseCase(sessionRepo)
	sessionUpdateUC := usecase.NewUpdateStudySessionUseCase(sessionRepo)
	sessionDeleteUC := usecase.NewDeleteStudySessionUseCase(sessionRepo)
	sessionHandler := handler.NewStudySessionHandler(sessionCreateUC, sessionListUC, sessionUpdateUC, sessionDeleteUC)

	// CourseSection
	sectionRepo := mongoRepo.NewCourseSectionRepository(db)
	sectionCreateUC := usecase.NewCreateCourseSectionUseCase(sectionRepo)
	sectionListUC := usecase.NewListCourseSectionsUseCase(sectionRepo)
	sectionGetUC := usecase.NewGetCourseSectionUseCase(sectionRepo)
	sectionUpdateUC := usecase.NewUpdateCourseSectionUseCase(sectionRepo)
	sectionDeleteUC := usecase.NewDeleteCourseSectionUseCase(sectionRepo)
	sectionHandler := handler.NewCourseSectionHandler(sectionCreateUC, sectionListUC, sectionGetUC, sectionUpdateUC, sectionDeleteUC)

	// CourseTopic
	topicRepo := mongoRepo.NewCourseTopicRepository(db)
	topicCreateUC := usecase.NewCreateCourseTopicUseCase(topicRepo)
	topicListUC := usecase.NewListCourseTopicsUseCase(topicRepo)
	topicGetUC := usecase.NewGetCourseTopicUseCase(topicRepo)
	topicUpdateUC := usecase.NewUpdateCourseTopicUseCase(topicRepo)
	topicDeleteUC := usecase.NewDeleteCourseTopicUseCase(topicRepo)
	topicHandler := handler.NewCourseTopicHandler(topicCreateUC, topicListUC, topicGetUC, topicUpdateUC, topicDeleteUC)

	// Project
	projectRepo := mongoRepo.NewProjectRepository(db)
	projectCreateUC := usecase.NewCreateProjectUseCase(projectRepo)
	projectListUC := usecase.NewListProjectsUseCase(projectRepo)
	projectGetUC := usecase.NewGetProjectUseCase(projectRepo)
	projectUpdateUC := usecase.NewUpdateProjectUseCase(projectRepo)
	projectDeleteUC := usecase.NewDeleteProjectUseCase(projectRepo)
	projectHandler := handler.NewProjectHandler(projectCreateUC, projectListUC, projectGetUC, projectUpdateUC, projectDeleteUC)

	// VidaCriativa
	vidaCriativaRepo := mongoRepo.NewVidaCriativaRepository(db)
	vidaCriativaCreateUC := usecase.NewCreateVidaCriativaUseCase(vidaCriativaRepo)
	vidaCriativaListUC := usecase.NewListVidaCriativaUseCase(vidaCriativaRepo)
	vidaCriativaGetUC := usecase.NewGetVidaCriativaUseCase(vidaCriativaRepo)
	vidaCriativaUpdateUC := usecase.NewUpdateVidaCriativaUseCase(vidaCriativaRepo)
	vidaCriativaDeleteUC := usecase.NewDeleteVidaCriativaUseCase(vidaCriativaRepo)
	vidaCriativaHandler := handler.NewVidaCriativaHandler(vidaCriativaCreateUC, vidaCriativaListUC, vidaCriativaGetUC, vidaCriativaUpdateUC, vidaCriativaDeleteUC)

	// FinancialRecord
	financialRepo := mongoRepo.NewFinancialRecordRepository(db)
	financialCreateUC := usecase.NewCreateFinancialRecordUseCase(financialRepo)
	financialListUC := usecase.NewListFinancialRecordsUseCase(financialRepo)
	financialGetUC := usecase.NewGetFinancialRecordUseCase(financialRepo)
	financialUpdateUC := usecase.NewUpdateFinancialRecordUseCase(financialRepo)
	financialDeleteUC := usecase.NewDeleteFinancialRecordUseCase(financialRepo)
	financialHandler := handler.NewFinancialRecordHandler(financialCreateUC, financialListUC, financialGetUC, financialUpdateUC, financialDeleteUC)

	// CulinaryCategory
	culCategoryRepo := mongoRepo.NewCulinaryCategoryRepository(db)
	culCategoryCreateUC := usecase.NewCreateCulinaryCategoryUseCase(culCategoryRepo)
	culCategoryListUC := usecase.NewListCulinaryCategoriesUseCase(culCategoryRepo)
	culCategoryGetUC := usecase.NewGetCulinaryCategoryUseCase(culCategoryRepo)
	culCategoryUpdateUC := usecase.NewUpdateCulinaryCategoryUseCase(culCategoryRepo)
	culCategoryDeleteUC := usecase.NewDeleteCulinaryCategoryUseCase(culCategoryRepo)
	culCategoryHandler := handler.NewCulinaryCategoryHandler(culCategoryCreateUC, culCategoryListUC, culCategoryGetUC, culCategoryUpdateUC, culCategoryDeleteUC)

	// CulinariaRecipe
	culinariaRepo := mongoRepo.NewCulinariaRecipeRepository(db)
	culinariaCreateUC := usecase.NewCreateCulinariaRecipeUseCase(culinariaRepo)
	culinariaListUC := usecase.NewListCulinariaRecipesUseCase(culinariaRepo)
	culinariaGetUC := usecase.NewGetCulinariaRecipeUseCase(culinariaRepo)
	culinariaUpdateUC := usecase.NewUpdateCulinariaRecipeUseCase(culinariaRepo)
	culinariaDeleteUC := usecase.NewDeleteCulinariaRecipeUseCase(culinariaRepo)
	culinariaHandler := handler.NewCulinariaRecipeHandler(culinariaCreateUC, culinariaListUC, culinariaGetUC, culinariaUpdateUC, culinariaDeleteUC)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	noteHandler.RegisterRoutes(mux)
	resourceHandler.RegisterRoutes(mux)
	quizHandler.RegisterRoutes(mux)
	sessionHandler.RegisterRoutes(mux)
	sectionHandler.RegisterRoutes(mux)
	topicHandler.RegisterRoutes(mux)
	projectHandler.RegisterRoutes(mux)
	vidaCriativaHandler.RegisterRoutes(mux)
	financialHandler.RegisterRoutes(mux)
	culCategoryHandler.RegisterRoutes(mux)
	culinariaHandler.RegisterRoutes(mux)

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
