package usecase_test

import (
	"context"
	"errors"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockRepo struct {
	items  map[string]*entity.StudyItem
	nextID int
	err    error
}

func newMockRepo() *mockRepo {
	return &mockRepo{items: make(map[string]*entity.StudyItem)}
}

func (m *mockRepo) Create(_ context.Context, item *entity.StudyItem) (*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	item.ID = primitive.NewObjectID()
	m.items[item.ID.Hex()] = item
	return item, nil
}

func (m *mockRepo) GetByID(_ context.Context, id string) (*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	item, ok := m.items[id]
	if !ok {
		return nil, errors.New("item não encontrado")
	}
	return item, nil
}

func (m *mockRepo) List(_ context.Context, filter repository.Filter) ([]*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	var result []*entity.StudyItem
	for _, item := range m.items {
		if filter.Section != "" && item.Section != filter.Section {
			continue
		}
		if filter.Topic != "" && item.Topic != filter.Topic {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (m *mockRepo) Update(_ context.Context, id string, item *entity.StudyItem) (*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.items[id] = item
	return item, nil
}

func (m *mockRepo) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.items, id)
	return nil
}
