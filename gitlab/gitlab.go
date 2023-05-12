package gitlab

import (
	"fmt"
	"io"
	"net/http"
)

type GitLab struct {
	Host  string
	API   string
	Token string
}

func (me *GitLab) Do(method string, uri string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, me.API+uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", me.Token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
