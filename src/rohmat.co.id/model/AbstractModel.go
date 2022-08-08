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

type MasterModel struct {
	ID       sql.NullInt64
	ParentID sql.NullInt64
	Code     sql.NullString
	Name     sql.NullString
}
