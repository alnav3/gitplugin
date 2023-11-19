package service

import (
	"encoding/json"
	"io"
	"log"
	"model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStructure(c echo.Context) error {
	repo := new(model.Repository)

	if err := c.Bind(repo); err != nil {
		return err
	}

	githubUrl := "https://api.github.com/repos/" + repo.Owner + "/" + repo.Repo + "/git/trees/" + repo.Branch + "?recursive=1"

	body := getRequest(githubUrl)

	var response model.GithubResponse

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, response.Tree)
}

func GetFile(c echo.Context) error {
	file := new(model.File)

	if err := c.Bind(file); err != nil {
		return err
	}

	githubUrl := "https://raw.githubusercontent.com/" + file.Owner + "/" + file.Repo + "/" + file.Branch + "/" + file.Path
	body := getRequest(githubUrl)

	return c.String(http.StatusOK, string(body))
}

func getRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	handle(err)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	handle(err)
	body, err := io.ReadAll(resp.Body)
	handle(err)
	defer resp.Body.Close()
	return body
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
