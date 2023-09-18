package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func LocalImage(c echo.Context) error {
	filePath := fmt.Sprintf("resources/img/%s", c.Param("filename"))

	_, err := os.Stat(filePath)
	if err != nil {
		return c.String(http.StatusNotFound, "File not found")
	}

	return c.File(filePath)
}

func RemoteImage(c echo.Context, fileURL string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.Response().WriteHeader(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			c.Response().Header().Add(key, value)
		}
	}

	_, err = c.Response().Write(body)
	if err != nil {
		return err
	}

	return nil
}

func GithubImage(c echo.Context) error {
	fileReqURL := fmt.Sprintf(
		"https://raw.githubusercontent.com/lanjoni/lpi4noobs/main/.github/%s",
		c.Param("filename"),
	)

	return RemoteImage(c, fileReqURL)
}

func ContentImage(c echo.Context) error {
	fileReqURL := fmt.Sprintf(
		"https://raw.githubusercontent.com/lanjoni/lpi4noobs/main/content/img/%s?raw=true",
		c.Param("filename"),
	)

	return RemoteImage(c, fileReqURL)
}
