package types

import "time"

type Note struct {
	Content string `json:"content"`
	To      string `json:"to"`
}

type UsersTable struct {
	Username  string
	Pin       string
	CreatedAt time.Time
}

type NotesTable struct {
	Id        string
	Content   string
	Username  string
	CreatedAt string // time.Time
}
