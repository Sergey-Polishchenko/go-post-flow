package resolvers

import (
	"context"
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers/mock"
)

func Test_mutationResolver_CreatePost(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockPost := &model.Post{ID: "1", Title: "New Post", Content: "This is a new post."}
	input := model.PostInput{Title: "New Post", Content: "This is a new post."}

	mockStorage.On("CreatePost", input).Return(mockPost, nil)

	r := &mutationResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	got, err := r.CreatePost(ctx, input)
	if err != nil {
		t.Errorf("mutationResolver.CreatePost() error = %v, wantErr %v", err, false)
		return
	}

	if !reflect.DeepEqual(got, mockPost) {
		t.Errorf("mutationResolver.CreatePost() = %v, want %v", got, mockPost)
	}

	mockStorage.AssertCalled(t, "CreatePost", input)
}

func Test_mutationResolver_CreateComment(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockComment := &model.Comment{ID: "1", Text: "New Comment", Post: &model.Post{ID: "post_1"}}
	input := model.CommentInput{
		Text:     "New Comment",
		PostID:   "post_1",
		AuthorID: "user_1",
	}

	// Настройка ожиданий
	mockStorage.On("CreateComment", input).Return(mockComment, nil)
	mockStorage.On("BroadcastComment", mockComment).Once()

	r := &mutationResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	got, err := r.CreateComment(ctx, input)
	if err != nil {
		t.Errorf("mutationResolver.CreateComment() error = %v, wantErr %v", err, false)
		return
	}

	if !reflect.DeepEqual(got, mockComment) {
		t.Errorf("mutationResolver.CreateComment() = %v, want %v", got, mockComment)
	}

	mockStorage.AssertCalled(t, "CreateComment", input)
	mockStorage.AssertCalled(t, "BroadcastComment", mockComment)
	mockStorage.AssertExpectations(t)
}

func Test_mutationResolver_CreateUser(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockUser := &model.User{ID: "1", Name: "New User"}
	input := model.UserInput{Name: "New User"}

	mockStorage.On("CreateUser", input).Return(mockUser, nil)

	r := &mutationResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	got, err := r.CreateUser(ctx, input)
	if err != nil {
		t.Errorf("mutationResolver.CreateUser() error = %v, wantErr %v", err, false)
		return
	}

	if !reflect.DeepEqual(got, mockUser) {
		t.Errorf("mutationResolver.CreateUser() = %v, want %v", got, mockUser)
	}

	mockStorage.AssertCalled(t, "CreateUser", input)
}
