package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model

	Content          string `json:"content"`
	ChecksWithScript bool   `json:"checks_with_script"`
	Expects          string `json:"expects"`

	ParentArticleRefer uint
	ParentArticle      Article `gorm:"foreignKey:ParentArticleRefer"`
}

// should be used only when loading from exercise list resource.
type ExerciseRoot struct {
	ArticleTitle string     `json:"article_title"`
	Exercises    []Exercise `json:"exercises"`
}

func ReadExercises(db *gorm.DB) ([]Exercise, error) {
	var exercises []Exercise
	err := db.Find(&exercises).Error
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func PopulateExercises(resource string, rootArticle *Article) ([]Exercise, error) {
	var exercisesRoots []ExerciseRoot
	err := PopulateFromResource(&exercisesRoots, resource)
	if err != nil {
		return nil, err
	}

	var finalExercises []Exercise
	for _, exerciseRoot := range exercisesRoots {
		parentArticle := rootArticle.FindByTitle(exerciseRoot.ArticleTitle)
		if parentArticle == nil {
			return nil, errors.Errorf(
				"failed to find parent article with title: %s\n",
				exerciseRoot.ArticleTitle,
			)
		}

		for i := range exerciseRoot.Exercises {
			exerciseRoot.Exercises[i].ParentArticleRefer = parentArticle.ID
			finalExercises = append(finalExercises, exerciseRoot.Exercises[i])
		}
	}

	return finalExercises, nil
}

func (exercise *Exercise) Write(db *gorm.DB) error {
	err := db.Create(exercise).Error
	return err
}

func SortExercisesByParentArticle(exercises []Exercise, parentArticle *Article) []Exercise {
	var sortedExercises []Exercise
	for i := range exercises {
		if exercises[i].ParentArticleRefer == parentArticle.ID {
			sortedExercises = append(sortedExercises, exercises[i])
		}
	}

	return sortedExercises
}
