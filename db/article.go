package db

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

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

type ArticleTree struct {
	Article
	Children []*ArticleTree `json:"children"`
}

func PopulateArticles(resource string) (*Article, error) {
	data, err := os.ReadFile(resource)
	if err != nil {
		return nil, err
	}

	rootArticle := &Article{}
	err = json.Unmarshal(data, rootArticle)
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
