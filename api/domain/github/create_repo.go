package github

// {
// 	"name": "go-tutorial",
// 	"description": "This is the go tutorial repo",
// 	"homepage": "https://github.com",
// 	"private": false,
// 	"has_issues": true,
// 	"has_projects": true,
// 	"has_wiki": true
//}

// CreateRepoRequest model for creating a new repo
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse model for repo creation success
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

// RepoOwner model for the repo owner
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// RepoPermissions permissions for the repo
type RepoPermissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}
