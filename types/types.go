package types

import "time"

type UsersTable struct {
	Username  string
	Pin       string
	CreatedAt time.Time
}

type NotesTable struct {
	Id        string
	Content   string
	Username  string
	CreatedAt time.Time
}

type Notes struct {
	Id        string
	Content   string
	Username  string
	CreatedAt string
}
