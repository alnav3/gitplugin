package model

type Repo struct {
	Owner  string "json:owner"
	Repo   string "json:repo"
	Branch string "json:branch"
}

type GithubResponse struct {
	Tree []TreeItem `json:"tree"`
}

type TreeItem struct {
	Path string `json:"path"`
	Type string `json:"type"`
}
