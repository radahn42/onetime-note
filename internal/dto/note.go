package dto

type CreateNoteRequest struct {
	Content    string `json:"content" validate:"required,min=1,max=2000"`
	TTLSeconds *int   `json:"ttl_seconds,omitempty" validate:"omitempty,gte=0"`
}

type CreateNoteResponse struct {
	ID  string `json:"id" validate:"uuid"`
	URL string `json:"url" validate:"relative_url"`
}

type GetNoteResponse struct {
	Content string `json:"content" validate:"required,min=1,max=2000"`
}
