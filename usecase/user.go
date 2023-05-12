package usecase

import (
	"encoding/json"
	"io"

	"github.com.br/rafaelbmateus/gitlab-dashboard/entity"
	"github.com/rs/zerolog/log"
)

func (me *Usecase) GetUsers() ([]*entity.User, error) {
	users := make([]*entity.User, 0)

	for _, u := range me.Config.Users {
		res, err := me.GitLab.GetUser(u)
		if err != nil {
			return users, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error().Msgf("error on read body %q", err)

			return nil, err
		}

		var user entity.User

		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Error().Msgf("error on unmarshal user %q", err)

			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
