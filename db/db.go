package db

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	articlesResource  string = "resources/artifacts/articles.json"
	exercisesResource string = "resources/artifacts/exercises.json"
)

type DB struct {
	Refer       *gorm.DB
	RootArticle *Article
	Exercises   []Exercise
	Sessions    []Session
}

func New(filePath string) (*DB, error) {
	_, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Article{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Exercise{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Session{})
	if err != nil {
		return nil, err
	}

	rootArticle, err := ReadArticles(db)
	if err != nil {
		return nil, err
	} else if rootArticle == nil {
		rootArticle, err = PopulateArticles(articlesResource)
		if err != nil {
			return nil, err
		}

		err = rootArticle.Write(db)
		if err != nil {
			return nil, err
		}
	}

	exercises, err := ReadExercises(db)
	if err != nil {
		return nil, err
	} else if len(exercises) <= 0 {
		exercises, err = PopulateExercises(exercisesResource, rootArticle)
		if err != nil {
			return nil, err
		}

		for _, exercise := range exercises {
			err = exercise.Write(db)
			if err != nil {
				return nil, err
			}
		}
	}

	sessions, err := ReadSessions(db)
	if err != nil {
		return nil, err
	} else if len(sessions) <= 0 {
		// create default session if none exist.
		defaultSession := Session{
			Username:   "default",
			ArticleID:  rootArticle.ID,
			ExerciseID: exercises[0].ID,
		}

		err = defaultSession.Write(db)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, defaultSession)
	}

	return &DB{
		Refer:       db,
		RootArticle: rootArticle,
		Exercises:   exercises,
		Sessions:    sessions,
	}, nil
}
