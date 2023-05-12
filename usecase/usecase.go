package usecase

import (
	"github.com.br/rafaelbmateus/gitlab-dashboard/config"
	"github.com.br/rafaelbmateus/gitlab-dashboard/gitlab"
)

type Usecase struct {
	Config *config.Config
	GitLab *gitlab.GitLab
}

func NewUsecase(cfg *config.Config, gl *gitlab.GitLab) *Usecase {
	return &Usecase{
		Config: cfg,
		GitLab: gl,
	}
}
