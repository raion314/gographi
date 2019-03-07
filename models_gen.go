// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gographi

type NewVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
	URL         string `json:"url"`
}

type Screenshot struct {
	ID      string `json:"id"`
	VideoID string `json:"videoId"`
	URL     string `json:"url"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Video struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	User        User          `json:"user"`
	URL         string        `json:"url"`
	CreatedAt   Timestamp     `json:"createdAt"`
	Screenshots []*Screenshot `json:"screenshots"`
	Related     []Video       `json:"related"`
}
