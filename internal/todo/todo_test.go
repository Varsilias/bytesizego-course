package todo

import (
	"context"
	"github.com/Varsilias/bytesizego-course/internal/db"
	"reflect"
	"strings"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	return m.items, nil
}

func (m *MockDB) SearchItem(ctx context.Context, searchString string) ([]db.Item, error) {
	var items = make([]db.Item, 0)
	for _, item := range m.items {
		if strings.Contains(strings.ToLower(item.Task), strings.ToLower(searchString)) {
			items = append(items, item)
		}
	}
	return items, nil
}

func (m *MockDB) GetItem(ctx context.Context, task string) *db.Item {
	for _, item := range m.items {
		if item.Task == task {
			return &item
		}
	}
	return nil
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name           string
		todosToAdd     []string
		searchString   string
		expectedResult []Item
	}{
		// TODO: Add test cases.
		{
			name:           "given a todo of \"Groceries Shopping\" and a searchString of \"Gro\" I should get \"Groceries Shopping\"",
			todosToAdd:     []string{"Groceries Shopping"},
			searchString:   "Gro",
			expectedResult: []Item{{Task: "Groceries Shopping", Status: "PENDING"}},
		},
		{
			name:           "should still return \"Groceries Shopping\" when case does not match",
			todosToAdd:     []string{"Groceries Shopping"},
			searchString:   "gro",
			expectedResult: []Item{{Task: "Groceries Shopping", Status: "PENDING"}},
		},
		{
			name:           "should return even with spaces",
			todosToAdd:     []string{"Groceries Shopping"},
			searchString:   "groceries",
			expectedResult: []Item{{Task: "Groceries Shopping", Status: "PENDING"}},
		},
		{
			name:           "should return even with spaces at the start of the todo item",
			todosToAdd:     []string{" Groceries Shopping"},
			searchString:   "groceries",
			expectedResult: []Item{{Task: " Groceries Shopping", Status: "PENDING"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := NewService(m)

			for _, toAdd := range tt.todosToAdd {
				if err := svc.Add(toAdd); err != nil {
					t.Error(err)
				}
			}
			got, err := svc.Search(tt.searchString)
			if err != nil {
				t.Errorf("Search() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}

func TestService_Add(t *testing.T) {

	tests := []struct {
		name           string
		todosToAdd     []string
		expectedResult []Item
	}{
		// TODO: Add test cases.
		{
			name: "Add todos successfully",
			todosToAdd: []string{
				"Learn Go",
				"Learn the standard library",
				"Go Shopping",
				"Get Groceries"},
			expectedResult: []Item{
				{Task: "Learn Go", Status: "PENDING"},
				{Task: "Learn the standard library", Status: "PENDING"},
				{Task: "Go Shopping", Status: "PENDING"},
				{Task: "Get Groceries", Status: "PENDING"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := NewService(m)

			for _, toAdd := range tt.todosToAdd {
				if err := svc.Add(toAdd); err != nil {
					t.Error(err)
				}
			}

			got, err := svc.GetAll()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Add() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
