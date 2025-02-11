package broadcast

import (
	"testing"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func TestNewBroadcast(t *testing.T) {
	b := NewBroadcast()
	if b == nil {
		t.Error("NewBroadcast() returned nil")
	}
	if b.CommentChannels == nil {
		t.Error("CommentChannels map is not initialized")
	}
}

func TestRegisterCommentChannel(t *testing.T) {
	b := NewBroadcast()
	postID := "post1"
	ch := make(chan *model.Comment)

	b.RegisterCommentChannel(postID, ch)

	if len(b.CommentChannels[postID]) != 1 {
		t.Errorf("Expected 1 channel, got %d", len(b.CommentChannels[postID]))
	}

	if b.CommentChannels[postID][0] != ch {
		t.Error("Channel was not registered correctly")
	}
}

func TestUnregisterCommentChannel(t *testing.T) {
	b := NewBroadcast()
	postID := "post1"
	ch := make(chan *model.Comment)

	b.RegisterCommentChannel(postID, ch)

	b.UnregisterCommentChannel(postID, ch)

	if len(b.CommentChannels[postID]) != 0 {
		t.Errorf("Expected 0 channels, got %d", len(b.CommentChannels[postID]))
	}
}

func TestBroadcastComment(t *testing.T) {
	b := NewBroadcast()
	postID := "post1"
	ch := make(chan *model.Comment, 1) // Буферизованный канал, чтобы не блокировать
	comment := &model.Comment{Post: &model.Post{ID: postID}}

	b.RegisterCommentChannel(postID, ch)

	b.BroadcastComment(comment)

	select {
	case receivedComment := <-ch:
		if receivedComment != comment {
			t.Error("Received comment does not match the sent comment")
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for comment")
	}
}

func TestBroadcastComment_NoChannels(t *testing.T) {
	b := NewBroadcast()
	comment := &model.Comment{Post: &model.Post{ID: "post1"}}
	b.BroadcastComment(comment)
}

func TestBroadcastComment_MultipleChannels(t *testing.T) {
	b := NewBroadcast()
	postID := "post1"
	ch1 := make(chan *model.Comment, 1)
	ch2 := make(chan *model.Comment, 1)
	comment := &model.Comment{Post: &model.Post{ID: postID}}

	b.RegisterCommentChannel(postID, ch1)
	b.RegisterCommentChannel(postID, ch2)

	b.BroadcastComment(comment)

	select {
	case receivedComment := <-ch1:
		if receivedComment != comment {
			t.Error("Received comment does not match the sent comment in ch1")
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for comment in ch1")
	}

	select {
	case receivedComment := <-ch2:
		if receivedComment != comment {
			t.Error("Received comment does not match the sent comment in ch2")
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for comment in ch2")
	}
}

func TestUnregisterCommentChannel_NonExistentChannel(t *testing.T) {
	b := NewBroadcast()
	postID := "post1"
	ch := make(chan *model.Comment)

	b.UnregisterCommentChannel(postID, ch)
}
