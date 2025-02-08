package mock

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetPost(id string) (*model.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockStorage) GetPosts(limit, offset *int) ([]*model.Post, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*model.Post), args.Error(1)
}

func (m *MockStorage) GetUser(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockStorage) GetComments(postID string, limit, offset *int) ([]*model.Comment, error) {
	args := m.Called(postID, limit, offset)
	return args.Get(0).([]*model.Comment), args.Error(1)
}

func (m *MockStorage) GetChildren(commentID string, limit, offset *int) ([]*model.Comment, error) {
	args := m.Called(commentID, limit, offset)
	return args.Get(0).([]*model.Comment), args.Error(1)
}

func (m *MockStorage) CreatePost(input model.PostInput) (*model.Post, error) {
	args := m.Called(input)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	args := m.Called(input)
	return args.Get(0).(*model.Comment), args.Error(1)
}

func (m *MockStorage) CreateUser(input model.UserInput) (*model.User, error) {
	args := m.Called(input)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockStorage) RegisterCommentChannel(postID string, ch chan *model.Comment) {
	m.Called(postID, ch)
}

func (m *MockStorage) UnregisterCommentChannel(postID string, ch chan *model.Comment) {
	m.Called(postID, ch)
}

func (m *MockStorage) BroadcastComment(comment *model.Comment) {
	m.Called(comment)
}
