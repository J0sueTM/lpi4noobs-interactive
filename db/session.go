package db

import "gorm.io/gorm"

type Session struct {
	gorm.Model

	Username string

	ArticleID uint
	Article   Article

	ExerciseID uint // FIXME
}

func ReadSessions(db *gorm.DB) ([]Session, error) {
	var sessions []Session
	err := db.Find(&sessions).Error
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func WriteSession(db *gorm.DB, session *Session) error {
	err := db.Create(session).Error
	return err
}
