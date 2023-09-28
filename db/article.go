package db

import (
	"errors"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	Title   string `json:"title"`
	Remote  string `json:"remote"`
	Content []byte

	ParentRefer uint
	Children    []Article `json:"children" gorm:"foreignKey:ParentRefer"`
}

func PopulateArticles(resource string) (*Article, error) {
	rootArticle := &Article{}
	err := PopulateFromResource(rootArticle, resource)
	if err != nil {
		return nil, err
	}

	err = rootArticle.FetchContent(true)
	if err != nil {
		return nil, err
	}

	return rootArticle, err
}

func (article *Article) FetchContent(recursive bool) error {
	log.Printf("retrieving %s\n", article.Remote)

	resp, err := http.Get(article.Remote)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	article.Content = body

	if !recursive {
		goto skip_recursion
	}

	for i := range article.Children {
		err = article.Children[i].FetchContent(true)
		if err != nil {
			return err
		}
	}

skip_recursion:
	return nil
}

func ReadArticles(db *gorm.DB) (*Article, error) {
	var articles []Article
	err := db.Where("title = 'Início'").Find(&articles).Error
	if err != nil {
		return nil, err
	} else if len(articles) <= 0 {
		return nil, nil
	}

	rootArticle := &articles[0]
	err = readAndPopulateArticleChildren(db, rootArticle)
	if err != nil {
		return nil, err
	}

	return rootArticle, nil
}

func readAndPopulateArticleChildren(db *gorm.DB, root *Article) error {
	var children []Article
	err := db.Where("parent_refer = ?", root.ID).Find(&children).Error
	if err != nil {
		return err
	} else if len(children) <= 0 {
		return nil
	}

	root.Children = children
	for i := range root.Children {
		readAndPopulateArticleChildren(db, &root.Children[i])
	}

	return nil
}

func (rootArticle *Article) Write(db *gorm.DB) error {
	res := db.Create(rootArticle)
	return res.Error
}

func (rootArticle *Article) FindByID(id uint) *Article {
	if rootArticle.ID == id {
		return rootArticle
	}

	for _, article := range rootArticle.Children {
		foundArticle := article.FindByID(id)
		if foundArticle != nil {
			return foundArticle
		}
	}

	return nil
}

func (rootArticle *Article) FindByTitle(title string) *Article {
	if rootArticle.Title == title {
		return rootArticle
	}

	for _, article := range rootArticle.Children {
		foundArticle := article.FindByTitle(title)
		if foundArticle != nil {
			return foundArticle
		}
	}

	return nil
}

func (rootArticle *Article) FindParent(child *Article) *Article {
	if rootArticle.ID == child.ID {
		return child
	}

	for i, article := range rootArticle.Children {
		foundChild := article.FindParent(child)
		// if child is returned, return again the parent instead.
		if foundChild != nil {
			if article.ID == child.ID {
				return rootArticle
			}

			return &rootArticle.Children[i]
		}
	}

	return nil
}

func (rootArticle *Article) FindNext(article *Article) (*Article, error) {
	if len(article.Children) > 0 {
		return &article.Children[0], nil
	}

	parentArticle := rootArticle.FindParent(article)
	if parentArticle == nil {
		return nil, errors.New("failed to find article's parent")
	}

	for i, child := range parentArticle.Children {
		if child.ID == article.ID && (i+1) < len(parentArticle.Children) {
			return &parentArticle.Children[i+1], nil
		}
	}

	// parent is last child, go back
	actualParent := rootArticle.FindParent(parentArticle)
	if actualParent != nil {
		for i, child := range actualParent.Children {
			if child.ID == parentArticle.ID && (i+1) < len(actualParent.Children) {
				return &actualParent.Children[i+1], nil
			}
		}
	}

	return nil, errors.New("failed to find next article")
}

func AssociateParentID(id uint) uint {
	if id >= 7 && id <= 10 {
		return 2
	} else if id >= 11 && id <= 15 {
		return 3
	} else if id >= 16 && id <= 21 {
		return 4
	} else if id >= 22 && id <= 24 {
		return 5
	} else if id >= 25 && id <= 28 {
		return 6
	}

	return 1
}
