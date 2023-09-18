package db

import "gorm.io/gorm"

type Session struct {
	gorm.Model

	Username string

	ArticleID int
	Article   Article

	ExerciseID int // FIXME
}
