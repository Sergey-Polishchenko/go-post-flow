package broadcast

import (
	"sync"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

type Broadcast struct {
	mu              sync.RWMutex
	CommentChannels map[string][]chan *model.Comment
}

func NewBroadcast() *Broadcast {
	return &Broadcast{
		CommentChannels: make(map[string][]chan *model.Comment),
	}
}

func (b *Broadcast) RegisterCommentChannel(postID string, ch chan *model.Comment) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.CommentChannels[postID] = append(b.CommentChannels[postID], ch)
}

func (b *Broadcast) UnregisterCommentChannel(postID string, ch chan *model.Comment) {
	b.mu.Lock()
	defer b.mu.Unlock()
	channels := b.CommentChannels[postID]
	for i, c := range channels {
		if c == ch {
			b.CommentChannels[postID] = append(channels[:i], channels[i+1:]...)
			return
		}
	}
}

func (b *Broadcast) BroadcastComment(comment *model.Comment) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, ch := range b.CommentChannels[comment.Post.ID] {
		select {
		case ch <- comment:
		default:
		}
	}
}
