package resolvers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers/mock"
	flowerrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func setupMutationResolver(t *testing.T) (*mutationResolver, *mock.MockStorage) {
	t.Helper()
	mockStorage := new(mock.MockStorage)
	return &mutationResolver{
		Resolver: &Resolver{
			storage: mockStorage,
			comLim:  1000,
		},
	}, mockStorage
}

func TestMutationResolver_CreatePost(t *testing.T) {
	r, mockStorage := setupMutationResolver(t)
	mockPost := &model.Post{ID: "1", Title: "New Post", Content: "This is a new post."}
	input := model.PostInput{Title: "New Post", Content: "This is a new post."}

	t.Run("Success", func(t *testing.T) {
		mockStorage.On("CreatePost", input).Return(mockPost, nil).Once()

		ctx := context.Background()
		got, err := r.CreatePost(ctx, input)
		require.NoError(t, err)
		assert.Equal(t, mockPost, got)

		mockStorage.AssertCalled(t, "CreatePost", input)
	})

	t.Run("Error", func(t *testing.T) {
		mockStorage.On("CreatePost", input).Return(nil, errors.New("storage error")).Once()

		ctx := context.Background()
		_, err := r.CreatePost(ctx, input)
		assert.Error(t, err)

		mockStorage.AssertCalled(t, "CreatePost", input)
	})
}

func TestMutationResolver_CreateComment(t *testing.T) {
	r, mockStorage := setupMutationResolver(t)
	mockComment := &model.Comment{ID: "1", Text: "New Comment", Post: &model.Post{ID: "post_1"}}
	input := model.CommentInput{
		Text:     "New Comment",
		PostID:   "post_1",
		AuthorID: "user_1",
	}

	t.Run("Success", func(t *testing.T) {
		mockStorage.On("CreateComment", input).Return(mockComment, nil).Once()
		mockStorage.On("BroadcastComment", mockComment).Once()

		ctx := context.Background()
		got, err := r.CreateComment(ctx, input)
		require.NoError(t, err)
		assert.Equal(t, mockComment, got)

		mockStorage.AssertCalled(t, "CreateComment", input)
		mockStorage.AssertCalled(t, "BroadcastComment", mockComment)
	})

	t.Run("CommentTooLong", func(t *testing.T) {
		longInput := model.CommentInput{
			Text:     "This is a very long comment that exceeds the limit",
			PostID:   "post_1",
			AuthorID: "user_1",
		}

		r.Resolver.comLim = 10

		ctx := context.Background()
		_, err := r.CreateComment(ctx, longInput)
		assert.True(t, errors.Is(err, flowerrors.ErrCommentTooLong))
	})

	t.Run("StorageError", func(t *testing.T) {
		mockStorage.On("CreateComment", input).Return(nil, errors.New("storage error")).Once()

		ctx := context.Background()
		_, err := r.CreateComment(ctx, input)
		assert.Error(t, err)

		mockStorage.AssertCalled(t, "CreateComment", input)
		mockStorage.AssertNotCalled(t, "BroadcastComment")
	})
}

func TestMutationResolver_CreateUser(t *testing.T) {
	r, mockStorage := setupMutationResolver(t)
	mockUser := &model.User{ID: "1", Name: "New User"}
	input := model.UserInput{Name: "New User"}

	t.Run("Success", func(t *testing.T) {
		mockStorage.On("CreateUser", input).Return(mockUser, nil).Once()

		ctx := context.Background()
		got, err := r.CreateUser(ctx, input)
		require.NoError(t, err)
		assert.Equal(t, mockUser, got)

		mockStorage.AssertCalled(t, "CreateUser", input)
	})

	t.Run("Error", func(t *testing.T) {
		mockStorage.On("CreateUser", input).Return(nil, errors.New("storage error")).Once()

		ctx := context.Background()
		_, err := r.CreateUser(ctx, input)
		assert.Error(t, err)

		mockStorage.AssertCalled(t, "CreateUser", input)
	})
}

func TestCommentResolver_Children(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	r := &commentResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	mockComments := []*model.Comment{
		{ID: "1", Text: "Child Comment 1"},
		{ID: "2", Text: "Child Comment 2"},
	}

	mockStorage.On("GetChildren", "parent-comment-id").Return(mockComments, nil)

	ctx := context.Background()
	parentComment := &model.Comment{ID: "parent-comment-id"}
	limit := 2
	offset := 0
	depth := 1
	expand := false

	got, err := r.Children(ctx, parentComment, &limit, &offset, &depth, &expand)
	require.NoError(t, err)
	assert.Equal(t, mockComments[offset:limit], got)

	mockStorage.AssertCalled(t, "GetChildren", "parent-comment-id")
}

func TestPostResolver_Comments(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	r := &postResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	mockComments := []*model.Comment{
		{ID: "1", Text: "Comment 1"},
		{ID: "2", Text: "Comment 2"},
	}

	mockStorage.On("GetComments", "post-id").Return(mockComments, nil)

	ctx := context.Background()
	post := &model.Post{ID: "post-id"}
	limit := 2
	offset := 0
	depth := 0
	expand := false

	got, err := r.Comments(ctx, post, &limit, &offset, &depth, &expand)
	require.NoError(t, err)
	assert.Equal(t, mockComments[offset:limit], got)

	mockStorage.AssertCalled(t, "GetComments", "post-id")
}

func TestQueryResolver_Post(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	r := &queryResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	mockPost := &model.Post{ID: "1", Title: "Test Post", Content: "This is a test post."}
	mockStorage.On("GetPost", "1").Return(mockPost, nil)

	ctx := context.Background()
	got, err := r.Post(ctx, "1")
	require.NoError(t, err)
	assert.Equal(t, mockPost, got)

	mockStorage.AssertCalled(t, "GetPost", "1")
}

func TestQueryResolver_Posts(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	r := &queryResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	mockPosts := []*model.Post{
		{ID: "1", Title: "Test Post 1", Content: "This is a test post 1."},
		{ID: "2", Title: "Test Post 2", Content: "This is a test post 2."},
	}

	limit := 10
	offset := 0

	mockStorage.On("GetPosts", &limit, &offset).Return(mockPosts, nil)

	ctx := context.Background()
	got, err := r.Posts(ctx, &limit, &offset)
	require.NoError(t, err)
	assert.Equal(t, mockPosts, got)

	mockStorage.AssertCalled(t, "GetPosts", &limit, &offset)
}

func TestSubscriptionResolver_CommentAdded(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	r := &subscriptionResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	postID := "1"
	expectedComment := &model.Comment{ID: "1", Text: "Test Comment"}

	var commentChan chan *model.Comment
	mockStorage.On("RegisterCommentChannel", postID, tmock.AnythingOfType("chan *model.Comment")).
		Run(func(args tmock.Arguments) {
			commentChan = args.Get(1).(chan *model.Comment)
		}).
		Return().
		Once()
	mockStorage.On("UnregisterCommentChannel", postID, tmock.AnythingOfType("chan *model.Comment")).
		Return().
		Once()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := r.CommentAdded(ctx, postID)
	require.NoError(t, err)

	go func() {
		commentChan <- expectedComment
	}()

	select {
	case receivedComment := <-ch:
		assert.Equal(t, expectedComment, receivedComment)
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for comment")
	}

	cancel()

	select {
	case _, ok := <-ch:
		assert.False(t, ok)
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for channel closure")
	}

	mockStorage.AssertExpectations(t)
}
