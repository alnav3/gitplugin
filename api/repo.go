package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Repository struct {
    Owner  string `json:"owner"`
    Repo   string `json:"repo"`
    Branch string `json:"branch"`
}

type GithubResponse struct {
    Tree []struct {
        Path string `json:"path"`
        Type string `json:"type"`
    } `json:"tree"`
}

func RepoHandler(w http.ResponseWriter, r *http.Request) {
    var repo Repository

    // Decodificar el cuerpo de la solicitud en la estructura Repository
    if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    githubUrl := "https://api.github.com/repos/" + repo.Owner + "/" + repo.Repo + "/git/trees/" + repo.Branch + "?recursive=1"

    body := getRequest(githubUrl)

    var response GithubResponse

    if err := json.Unmarshal(body, &response); err != nil {
        log.Println(err)
        http.Error(w, "Error procesando la respuesta de GitHub", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response.Tree)
}

func getRequest(url string) []byte {
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

