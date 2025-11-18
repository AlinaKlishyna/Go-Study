package core

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []Meetup  `json:"meetups"`
}

type Meetup struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Note     string `json:"note"`
	User     *User  `json:"user"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
