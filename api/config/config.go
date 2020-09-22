package config

import "os"

// LogLevel indicates the details level of the logging messages
const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	goEnvironment        = "GO_ENVIRONMENT"
	production           = "production"

	LogLevel = "info"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// GetGithubAccessToken -
func GetGithubAccessToken() string {
	return githubAccessToken
}

// IsProduction -
func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
