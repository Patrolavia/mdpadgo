package model

// User record
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"-"`
	Image string `json:"image"`
}