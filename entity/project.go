package entity

import "time"

type Project struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         string    `json:"created_at"`
	Default_branch    string    `json:"default_branch"`
	Topics            []string  `json:"topics"`
	WebURL            string    `json:"web_url"`
	Forks_count       int       `json:"forks_count"`
	StarCount         int       `json:"star_count"`
	LastActivityAt    time.Time `json:"last_activity_at"`
	Archived          bool      `json:"archived"`
	Visibility        string    `json:"visibility"`
	UpdatedAt         time.Time `json:"updated_at"`
	Commits           []Commit  `json:"commits"`
	LastCommit        Commit    `json:"last_commit"`
}
