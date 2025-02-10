package inmemory

import "github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"

func (s *InMemoryStorage) RegisterCommentChannel(postID string, ch chan *model.Comment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.commentChannels[postID] = append(s.commentChannels[postID], ch)
}

func (s *InMemoryStorage) UnregisterCommentChannel(postID string, ch chan *model.Comment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	channels := s.commentChannels[postID]
	for i, c := range channels {
		if c == ch {
			s.commentChannels[postID] = append(channels[:i], channels[i+1:]...)
			return
		}
	}
}

func (s *InMemoryStorage) BroadcastComment(comment *model.Comment) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, ch := range s.commentChannels[comment.Post.ID] {
		select {
		case ch <- comment:
		default:
		}
	}
}
