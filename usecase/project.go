package usecase

import (
	"encoding/json"
	"io"

	"github.com.br/rafaelbmateus/gitlab-dashboard/entity"
	"github.com/rs/zerolog/log"
)

func (me *Usecase) GetProjects() ([]*entity.Project, error) {
	log.Debug().Msg("get projects")
	projects := make([]*entity.Project, 0)

	for _, p := range me.Config.Projects {
		log.Debug().Msgf("get project %s", p)
		res, err := me.GitLab.GetProject(p)
		if err != nil {
			log.Error().Msgf("error on get projects %q", err)

			return projects, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error().Msgf("error on read body %q", err)

			return nil, err
		}

		var project entity.Project

		err = json.Unmarshal(body, &project)
		if err != nil {
			log.Error().Msgf("error on unmarshal project %q", err)

			return nil, err
		}

		commits, err := me.GetProjectCommits(p)
		project.Commits = commits
		if err != nil {
			log.Error().Msgf("error on get project commits %q", err)
		}

		project.LastCommit = commits[0]

		projects = append(projects, &project)
	}

	return projects, nil
}
