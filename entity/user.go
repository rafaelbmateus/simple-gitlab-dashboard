package entity

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Bio          string `json:"bio"`
	JobTitle     string `json:"job_title"`
	Location     string `json:"location"`
	State        string `json:"state"`
	WebURL       string `json:"web_url"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
	Organization string `json:"organization"`
	CreatedAt    string `json:"created_at"`
}
