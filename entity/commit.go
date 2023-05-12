package entity

import "time"

type Commit struct {
	ID            string    `json:"id"`
	ShortID       string    `json:"short_id"`
	Title         string    `json:"title"`
	Message       string    `json:"message"`
	AuthorName    string    `json:"author_name"`
	AuthorEmail   string    `json:"author_email"`
	WebURL        string    `json:"web_url"`
	CommittedDate time.Time `json:"committed_date"`
	CreatedAt     time.Time `json:"created_at"`
}
