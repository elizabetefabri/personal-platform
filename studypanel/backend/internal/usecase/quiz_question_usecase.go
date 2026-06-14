package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// --- Create ---

type CreateQuizQuestionInput struct {
	Section     string                 `json:"section"`
	Topic       string                 `json:"topic"`
	Question    string                 `json:"question"`
	Answer      string                 `json:"answer"`
	Explanation string                 `json:"explanation"`
	Difficulty  entity.QuizDifficulty  `json:"difficulty"`
	Tags        []string               `json:"tags"`
}

type CreateQuizQuestionUseCase struct {
	repo repository.QuizQuestionRepository
}

func NewCreateQuizQuestionUseCase(repo repository.QuizQuestionRepository) *CreateQuizQuestionUseCase {
	return &CreateQuizQuestionUseCase{repo: repo}
}

func (uc *CreateQuizQuestionUseCase) Execute(ctx context.Context, input CreateQuizQuestionInput) (*entity.QuizQuestion, error) {
	if input.Section == "" {
		return nil, errors.New("section é obrigatório")
	}
	if input.Topic == "" {
		return nil, errors.New("topic é obrigatório")
	}
	if input.Question == "" {
		return nil, errors.New("question é obrigatório")
	}
	if input.Answer == "" {
		return nil, errors.New("answer é obrigatório")
	}
	if !entity.ValidQuizDifficulties[input.Difficulty] {
		return nil, errors.New("difficulty inválido")
	}

	tags := input.Tags
	if tags == nil {
		tags = []string{}
	}

	question := &entity.QuizQuestion{
		Section:     input.Section,
		Topic:       input.Topic,
		Question:    input.Question,
		Answer:      input.Answer,
		Explanation: input.Explanation,
		Difficulty:  input.Difficulty,
		Tags:        tags,
		CreatedAt:   time.Now().UTC(),
	}

	return uc.repo.Create(ctx, question)
}

// --- List ---

type ListQuizQuestionsUseCase struct {
	repo repository.QuizQuestionRepository
}

func NewListQuizQuestionsUseCase(repo repository.QuizQuestionRepository) *ListQuizQuestionsUseCase {
	return &ListQuizQuestionsUseCase{repo: repo}
}

func (uc *ListQuizQuestionsUseCase) Execute(ctx context.Context, section, topic string) ([]*entity.QuizQuestion, error) {
	return uc.repo.ListByTopic(ctx, section, topic)
}

// --- Update ---

type UpdateQuizQuestionInput struct {
	Section     string                `json:"section"`
	Topic       string                `json:"topic"`
	Question    string                `json:"question"`
	Answer      string                `json:"answer"`
	Explanation string                `json:"explanation"`
	Difficulty  entity.QuizDifficulty `json:"difficulty"`
	Tags        []string              `json:"tags"`
}

type UpdateQuizQuestionUseCase struct {
	repo repository.QuizQuestionRepository
}

func NewUpdateQuizQuestionUseCase(repo repository.QuizQuestionRepository) *UpdateQuizQuestionUseCase {
	return &UpdateQuizQuestionUseCase{repo: repo}
}

func (uc *UpdateQuizQuestionUseCase) Execute(ctx context.Context, id string, input UpdateQuizQuestionInput) (*entity.QuizQuestion, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if input.Question == "" {
		return nil, errors.New("question é obrigatório")
	}
	if input.Answer == "" {
		return nil, errors.New("answer é obrigatório")
	}
	if !entity.ValidQuizDifficulties[input.Difficulty] {
		return nil, errors.New("difficulty inválido")
	}

	tags := input.Tags
	if tags == nil {
		tags = []string{}
	}

	question := &entity.QuizQuestion{
		Section:     input.Section,
		Topic:       input.Topic,
		Question:    input.Question,
		Answer:      input.Answer,
		Explanation: input.Explanation,
		Difficulty:  input.Difficulty,
		Tags:        tags,
	}

	return uc.repo.Update(ctx, id, question)
}

// --- Delete ---

type DeleteQuizQuestionUseCase struct {
	repo repository.QuizQuestionRepository
}

func NewDeleteQuizQuestionUseCase(repo repository.QuizQuestionRepository) *DeleteQuizQuestionUseCase {
	return &DeleteQuizQuestionUseCase{repo: repo}
}

func (uc *DeleteQuizQuestionUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
