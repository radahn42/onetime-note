package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/radahn42/onetime-note/internal/dto"
	"github.com/radahn42/onetime-note/internal/service"
)

type NoteHandler struct {
	svc *service.NoteService
}

func NewNoteHandler(svc *service.NoteService) *NoteHandler {
	return &NoteHandler{svc: svc}
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.svc.Create(r.Context(), req.Content, req.TTLSeconds)
	if err != nil {
		http.Error(w, "failed to create note", http.StatusInternalServerError)
		return
	}

	resp := dto.CreateNoteResponse{
		ID:  id,
		URL: "/api/notes/" + id,
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *NoteHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	content, err := h.svc.Get(r.Context(), id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	resp := dto.GetNoteResponse{
		Content: content,
	}
	json.NewEncoder(w).Encode(resp)
}
