package usecase

import (
	"encoding/json"
	"io"

	"github.com.br/rafaelbmateus/gitlab-dashboard/entity"
	"github.com/rs/zerolog/log"
)

func (me *Usecase) GetProjectCommits(id string) ([]entity.Commit, error) {
	res, err := me.GitLab.GetProjectCommits(id)
	if err != nil {
		log.Error().Msgf("error on get project last commit %q", err)

		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Msgf("error on read body %q", err)

		return nil, err
	}

	var commits []entity.Commit

	err = json.Unmarshal(body, &commits)
	if err != nil {
		log.Error().Msgf("error on unmarshal commits %q", err)

		return nil, err
	}

	return commits, nil
}
