package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Varsilias/bytesizego-course/internal/db"
	"log"
)

type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

type Manager interface {
	InsertItem(ctx context.Context, item db.Item) error
	GetAllItems(ctx context.Context) ([]db.Item, error)
	SearchItem(ctx context.Context, searchString string) ([]db.Item, error)
	GetItem(ctx context.Context, task string) *db.Item
}

type Service struct {
	db Manager
}

func NewService(db Manager) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) Add(todo string) error {
	todoExists := svc.db.GetItem(context.Background(), todo)

	if todoExists != nil {
		return errors.New("task already exists")
	}

	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "PENDING",
	}); err != nil {
		return fmt.Errorf("error inserting todo: %w", err)
	}

	return nil
}

func (svc *Service) GetAll() ([]Item, error) {
	var result = make([]Item, 0)
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}

	for _, item := range items {
		result = append(result, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}

	return result, nil
}

func (svc *Service) Search(searchString string) ([]Item, error) {
	var result = make([]Item, 0)
	items, err := svc.db.SearchItem(context.Background(), searchString)

	if err != nil {
		return nil, fmt.Errorf("failed to search items: %w", err)
	}

	for _, item := range items {
		result = append(result, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}

	return result, nil
}

func (svc *Service) todoExists(search string) bool {
	items, err := svc.db.GetAllItems(context.Background())

	if err != nil {
		log.Println("failed to get all items: ", err)
		return false
	}
	for _, item := range items {
		if search == item.Task {
			return true
		}
	}

	return false
}
