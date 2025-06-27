package dto

type CreateNoteRequest struct {
	Content    string `json:"content"`
	TTLSeconds int    `json:"ttl_seconds"`
}

type CreateNoteResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type GetNoteResponse struct {
	Content string `json:"content"`
}
