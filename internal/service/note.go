package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/radahn42/onetime-note/internal/storage"
	"time"
)

var (
	ErrNoteNotFound = errors.New("note not found")
	ErrInvalidTTL   = errors.New("invalid ttl")
)

type NoteService struct {
	store storage.Storage
}

func NewNoteService(store storage.Storage) *NoteService {
	return &NoteService{store: store}
}

func (s *NoteService) Create(ctx context.Context, content string, ttlSeconds int) (string, error) {
	if ttlSeconds <= 0 {
		return "", ErrInvalidTTL
	}

	id := uuid.NewString()
	err := s.store.Set(ctx, id, content, time.Duration(ttlSeconds)*time.Second)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *NoteService) Get(ctx context.Context, id string) (string, error) {
	content, err := s.store.Get(ctx, id)
	if err != nil {
		return "", ErrNoteNotFound
	}

	err = s.store.Delete(ctx, id)
	if err != nil {
		return "", ErrNoteNotFound
	}

	return content, nil
}
