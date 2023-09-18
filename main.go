package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/donseba/go-htmx"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/j0suetm/lpi4noobs-interactive/db"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	_, err := db.New("session.db")
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
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/bundle.css", func(c echo.Context) error {
		return c.File("views/style/bundle.css")
	})
	e.GET("/htmx.min.js", func(c echo.Context) error {
		return c.File("views/script/htmx.min.js")
	})
	e.GET("/img/:filename", func(c echo.Context) error {
		filePath := fmt.Sprintf("resources/img/%s", c.Param("filename"))

		_, err := os.Stat(filePath)
		if err != nil {
			return c.String(http.StatusNotFound, "File not found")
		}

		return c.File(filePath)
	})

	// sessionStt, err := session.New("cache/session.db")
	// if err != nil {
	//  	log.Fatal(err)
	// }
	//
	// sessionGroup := e.Group("/session")
	// {
	//  	sessionGroup.GET("", func(c echo.Context) error {
	//  		return c.Render(http.StatusOK, "session.html", sessionStt)
	//  	})
	//  	sessionGroup.GET("/content", sessionStt.Content)
	// }

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
