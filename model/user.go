package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// User represent a user.
type User struct {
	ModelDB
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Pass      string `json:"pass"`
}

func (u *User) BeforeSave() (err error) {
	u.Pass = u.HashPass(u.Pass)
	return nil
}

func (u *User) PassIsValid(pass string) bool {
	if u.Pass == u.HashPass(pass) {
		return true
	}
	return false
}

func (u User) MarshalJSON() ([]byte, error) {
	payload := make(map[string]interface{})
	payload["uuid"] = u.UUID
	payload["first_name"] = u.FirstName
	payload["last_name"] = u.LastName
	payload["email"] = u.Email
	return json.Marshal(payload)
}

func (u User) HashPass(pass string) string {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	return fmt.Sprintf("%x", h.Sum(nil))
}
