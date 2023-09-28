package api

import "github.com/j0suetm/lpi4noobs-interactive/db"

type State struct {
	Article   *db.Article
	Exercises []db.Exercise
	Session   *db.Session
}

type API struct {
	DB    *db.DB
	State *State
}
