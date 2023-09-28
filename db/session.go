package db

import "gorm.io/gorm"

type Session struct {
	gorm.Model

	Username string

	ArticleID uint
	Article   Article

	ExerciseID uint
	Exercise   Exercise
}

func ReadSessions(db *gorm.DB) ([]Session, error) {
	var sessions []Session
	err := db.Find(&sessions).Error
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (session *Session) Write(db *gorm.DB) error {
	err := db.Create(session).Error
	return err
}
