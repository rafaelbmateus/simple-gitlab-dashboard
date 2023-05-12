package main

import (
	"log"
	"os"

	"github.com.br/rafaelbmateus/gitlab-dashboard/config"
	"github.com.br/rafaelbmateus/gitlab-dashboard/gitlab"
	"github.com.br/rafaelbmateus/gitlab-dashboard/http"
	"github.com.br/rafaelbmateus/gitlab-dashboard/usecase"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("error on load config file %q", err)
	}

	gl := &gitlab.GitLab{
		Host:  cfg.GitLab.Host,
		API:   cfg.GitLab.API,
		Token: os.Getenv("GITLAB_TOKEN"),
	}

	usecase := usecase.NewUsecase(cfg, gl)

	http.NewHandler(usecase).Server()
}
