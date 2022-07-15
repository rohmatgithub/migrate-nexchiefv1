package model

import "database/sql"

type NexchiefAccount struct {
	ID     sql.NullInt64
	Schema sql.NullString
}

type ErrorModel struct {
	Error error
	Code  int
}
