package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


type File struct {
    Owner  string `json:"owner"`
    Repo   string `json:"repo"`
    Branch string `json:"branch"`
	Path string `json:"path"`
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
    file := new(File)

    if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	githubUrl := "https://raw.githubusercontent.com/" + file.Owner + "/" + file.Repo + "/" + file.Branch + "/" + file.Path
	body := getFile(githubUrl)

	w.Write([]byte(body))

}

func getFile(url string) []byte {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Add("Accept", "application/vnd.github.v3+json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    return body
}
