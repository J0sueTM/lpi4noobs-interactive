package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/donseba/go-htmx"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/j0suetm/lpi4noobs-interactive/api"
	"github.com/j0suetm/lpi4noobs-interactive/db"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	lpiDB, err := db.New("session.db")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(HtmxMiddleware)

	e.Renderer = echoview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		DisableCache: true,
		Funcs: template.FuncMap{
			"inc": func(i uint) uint {
				return i + 1
			},
			"dec": func(i uint) uint {
				return i - 1
			},
			"assocPar": db.AssociateParentID,
		},
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/bundle.css", func(c echo.Context) error {
		return c.File("views/style/bundle.css")
	})
	// yes i know i'm repeating myself
	e.GET("/htmx.min.js", func(c echo.Context) error {
		return c.File("views/script/htmx.min.js")
	})
	e.GET("/hyperscript.min.js", func(c echo.Context) error {
		return c.File("views/script/hyperscript.min.js")
	})

	imgG := e.Group("/img")
	{
		imgG.GET("/:filename", api.LocalImage)
		imgG.GET("/github/:filename", api.GithubImage)
		imgG.GET("/content/:filename", api.ContentImage)
	}

	lpiAPI := &api.API{
		DB:    lpiDB,
		State: &api.State{},
	}

	sessionG := e.Group("/session")
	{
		sessionG.GET("", lpiAPI.Session)
		sessionG.GET("/content", lpiAPI.Content)
		sessionG.GET("/content/exercises", lpiAPI.Exercises)
	}

	err = e.Start(":4192")
	e.Logger.Fatal(err)
}

func HtmxMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		hxh := htmx.HxRequestHeader{
			HxBoosted:               htmx.HxStrToBool(c.Request().Header.Get("HX-Boosted")),
			HxCurrentURL:            c.Request().Header.Get("HX-Current-URL"),
			HxHistoryRestoreRequest: htmx.HxStrToBool(c.Request().Header.Get("HX-History-Restore-Request")),
			HxPrompt:                c.Request().Header.Get("HX-Prompt"),
			HxRequest:               htmx.HxStrToBool(c.Request().Header.Get("HX-Request")),
			HxTarget:                c.Request().Header.Get("HX-Target"),
			HxTriggerName:           c.Request().Header.Get("HX-Trigger-Name"),
			HxTrigger:               c.Request().Header.Get("HX-Trigger"),
		}

		ctx = context.WithValue(ctx, htmx.ContextRequestHeader, hxh)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
