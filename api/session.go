package api

import (
	"bytes"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/j0suetm/lpi4noobs-interactive/db"
	"github.com/labstack/echo/v4"
)

func (lpiAPI *API) Session(c echo.Context) error {
	rootArticle := lpiAPI.DB.RootArticle
	session := &lpiAPI.DB.Sessions[0]

	article := rootArticle.FindByID(session.ArticleID)
	if article == nil {
		article = lpiAPI.DB.RootArticle
		lpiAPI.DB.Sessions[0].ArticleID = article.ID
	}

	var err error
	switch c.QueryParam("action") {
	case "next":
		article, err = rootArticle.FindNext(article)
		break
	case "prev":
		article, err = rootArticle.FindPrevious(article)
		break
	default:
		break
	}

	if err != nil {
		return err
	}

	session.ArticleID = article.ID

	lpiAPI.State.Article = article
	lpiAPI.State.Exercises = db.SortExercisesByParentArticle(
		lpiAPI.DB.Exercises,
		lpiAPI.State.Article,
	)

	return c.Render(http.StatusOK, "session.html", lpiAPI.State)
}

func (lpiAPI *API) Content(c echo.Context) error {
	content := lpiAPI.State.Article.Content
	content = bytes.ReplaceAll(content, []byte("content/img/"), []byte("/img/content/"))
	content = bytes.ReplaceAll(content, []byte("./.github/"), []byte("/img/github/"))

	extensions := parser.CommonExtensions |
		parser.AutoHeadingIDs |
		parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(content)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
	renderer := html.NewRenderer(opts)
	rndrDoc := markdown.Render(doc, renderer)

	return c.HTML(http.StatusOK, string(rndrDoc))
}

func (lpiAPI *API) Exercises(c echo.Context) error {
	return c.Render(http.StatusOK, "exercises.html", lpiAPI.State.Exercises)
}
