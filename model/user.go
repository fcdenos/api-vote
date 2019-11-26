package model

// User represent a user.
type User struct {
	ModelDB
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
