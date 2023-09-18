package api

import (
	"errors"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/labstack/echo/v4"
)

func (lpiAPI *API) Session(c echo.Context) error {
	lpiAPI.State.Article = lpiAPI.DB.RootArticle.FindByID(lpiAPI.DB.Sessions[0].ArticleID)
	if lpiAPI.State.Article == nil {
		return errors.New("failed to find current session's article")
	}

	return c.Render(http.StatusOK, "session.html", lpiAPI.State)
}

func (lpiAPI *API) Content(c echo.Context) error {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(lpiAPI.DB.RootArticle.Content)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
	renderer := html.NewRenderer(opts)
	rndrDoc := markdown.Render(doc, renderer)

	return c.HTML(http.StatusOK, string(rndrDoc))
}
