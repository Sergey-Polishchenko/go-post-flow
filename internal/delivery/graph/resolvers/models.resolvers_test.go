package resolvers

import (
	"context"
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers/mock"
)

func Test_commentResolver_Children(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockComments := []*model.Comment{
		{ID: "1", Text: "Child Comment 1"},
		{ID: "2", Text: "Child Comment 2"},
	}

	limit := 10
	offset := 0

	// Настраиваем мок
	mockStorage.On("GetChildren", "parent-comment-id", &limit, &offset).Return(mockComments, nil)

	r := &commentResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	parentComment := &model.Comment{ID: "parent-comment-id"}
	got, err := r.Children(ctx, parentComment, &limit, &offset)
	// Проверяем, что ошибки нет
	if err != nil {
		t.Errorf("commentResolver.Children() error = %v, wantErr %v", err, false)
		return
	}

	// Проверяем, что результат совпадает с ожидаемым
	if !reflect.DeepEqual(got, mockComments) {
		t.Errorf("commentResolver.Children() = %v, want %v", got, mockComments)
	}

	// Проверяем, что мок был вызван с правильными аргументами
	mockStorage.AssertCalled(t, "GetChildren", "parent-comment-id", &limit, &offset)
}

func Test_postResolver_Comments(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockComments := []*model.Comment{
		{ID: "1", Text: "Comment 1"},
		{ID: "2", Text: "Comment 2"},
	}

	limit := 10
	offset := 0

	// Настраиваем мок
	mockStorage.On("GetComments", "post-id", &limit, &offset).Return(mockComments, nil)

	r := &postResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	post := &model.Post{ID: "post-id"}
	got, err := r.Comments(ctx, post, &limit, &offset)
	// Проверяем, что ошибки нет
	if err != nil {
		t.Errorf("postResolver.Comments() error = %v, wantErr %v", err, false)
		return
	}

	// Проверяем, что результат совпадает с ожидаемым
	if !reflect.DeepEqual(got, mockComments) {
		t.Errorf("postResolver.Comments() = %v, want %v", got, mockComments)
	}

	// Проверяем, что мок был вызван с правильными аргументами
	mockStorage.AssertCalled(t, "GetComments", "post-id", &limit, &offset)
}
