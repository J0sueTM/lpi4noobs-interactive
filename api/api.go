package api

import "github.com/j0suetm/lpi4noobs-interactive/db"

type APIState struct {
	Article *db.Article
	// Exercise db.Exercise
	Session *db.Session
}

type API struct {
	DB    *db.DB
	State APIState
}
