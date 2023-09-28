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
		// FIXME
		article, err = rootArticle.FindNext(article)
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

	// parentArticleID, err := strconv.ParseInt(c.Param("parent-id"), 10, 64)
	// if err != nil {
	//  	return err
	// } else if lpiAPI.DB.RootArticle.ID > uint(parentArticleID) {
	//  	parentArticleID = int64(lpiAPI.DB.RootArticle.ID)
	// }
	//
	// childArticleID, err := strconv.ParseInt(c.Param("child-id"), 10, 64)
	// if err != nil {
	//  	return err
	// }
	//
	// var parentArticle *db.Article
	// var childArticle *db.Article
	//
	// parentArticle = lpiAPI.DB.RootArticle.FindByID(uint(parentArticleID))
	// if parentArticle == nil {
	//  	return errors.New("failed to find current session's parent article")
	// }
	//
	// if int(childArticleID) <= len(parentArticle.Children) && childArticleID >= 1 {
	//  	childArticle = &parentArticle.Children[childArticleID-1]
	// } else {
	//  	if childArticleID > 0 {
	//  		parentArticleID++
	//  	}
	//  	parentArticle = lpiAPI.DB.RootArticle.FindByID(uint(parentArticleID))
	//  	childArticle = parentArticle
	//  	childArticleID = 0
	// }
	//
	// // parentArticle := lpiAPI.DB.RootArticle
	//
	// // load from db instead
	// // if parentArticleID == 0 {
	// // 	childArticleID = int64(lpiAPI.DB.Sessions[0].ArticleID)
	// // }
	//
	// // parentArticle = lpiAPI.DB.RootArticle.FindByID(uint(parentArticleID))
	// // if parentArticle == nil {
	// // 	return errors.New("failed to find current session's parent article")
	// // }
	//
	// // if len(parentArticle.Children) < int(childArticleID) {
	// // 	parentArticle = lpiAPI.DB.RootArticle.FindByID(db.AssociateParentID(uint(childArticleID)))
	// // 	if parentArticle == nil {
	// // 		return errors.New(
	// // 			"failed to find current session's parent article associated to child",
	// // 		)
	// // 	}
	//
	// // 	childArticleID = int64(parentArticle.Children[0].ID + uint(childArticleID))
	// // }
	//
	// // skip_parent:
	// // 	childArticle := lpiAPI.DB.RootArticle
	// // 	if childArticleID != 0 {
	// // 		childArticle = childArticle.FindByID(uint(childArticleID))
	// // 		if childArticle == nil {
	// // 			return errors.New("failed to find current session's child article")
	// // 		}
	// // 	}
	//
	// lpiAPI.State.ParentArticleID = uint(parentArticleID)
	// lpiAPI.State.ChildArticleID = uint(childArticleID)
	// lpiAPI.State.Article = childArticle
	// lpiAPI.State.Exercises = db.SortExercisesByParentArticle(
	//  	lpiAPI.DB.Exercises,
	//  	lpiAPI.State.Article,
	// )
	//
	// return c.Render(http.StatusOK, "session.html", lpiAPI.State)
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
