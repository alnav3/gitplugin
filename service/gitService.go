package service

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Repo struct {
	Owner  string "json:owner"
	Repo   string "json:repo"
	Branch string "json:branch"
}

func GetStructure(c echo.Context) error {
	repo := new(Repo)

	if err := c.Bind(repo); err != nil {
		return err
	}

	githubUrl := "https://api.github.com/repos/" + repo.Owner + "/" + repo.Repo + "/git/trees/" + repo.Branch + "?recursive=1"

	req, err := http.NewRequest("GET", githubUrl, nil)
	handle(err)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	handle(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handle(err)
	return c.String(http.StatusOK, string(body))
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
