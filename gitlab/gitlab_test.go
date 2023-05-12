package gitlab_test

import (
	"os"

	"github.com.br/rafaelbmateus/gitlab-dashboard/gitlab"
)

func NewGitlab() *gitlab.GitLab {
	return &gitlab.GitLab{
		Host:  "https://gitlab.com/api/v4",
		Token: os.Getenv("GITLAB_TOKEN"),
	}
}
