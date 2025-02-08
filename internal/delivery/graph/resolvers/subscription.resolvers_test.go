package resolvers

import (
	"context"
	"testing"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	storageMock "github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscriptionResolver_CommentAdded(t *testing.T) {
	mockStorage := new(storageMock.MockStorage)
	postID := "1"
	expectedComment := &model.Comment{ID: "1", Text: "Test Comment"}

	var commentChan chan *model.Comment
	mockStorage.On("RegisterCommentChannel", postID, mock.AnythingOfType("chan *model.Comment")).
		Run(func(args mock.Arguments) {
			commentChan = args.Get(1).(chan *model.Comment)
		}).
		Return().
		Once()
	mockStorage.On("UnregisterCommentChannel", postID, mock.AnythingOfType("chan *model.Comment")).
		Return().
		Once()

	resolver := &subscriptionResolver{
		Resolver: &Resolver{storage: mockStorage},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := resolver.CommentAdded(ctx, postID)
	assert.NoError(t, err, "CommentAdded should not return an error")

	go func() {
		commentChan <- expectedComment
	}()

	select {
	case receivedComment := <-ch:
		assert.Equal(t, expectedComment, receivedComment, "Received comment should match expected")
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for comment")
	}

	cancel()

	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed after context cancellation")
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for channel closure")
	}

	mockStorage.AssertExpectations(t)
}
