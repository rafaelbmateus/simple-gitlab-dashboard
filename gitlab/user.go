package gitlab

import "net/http"

func (me *GitLab) GetCurrentUser(id string) (*http.Response, error) {
	return me.Do("GET", "/user", nil)
}

func (me *GitLab) GetUser(id string) (*http.Response, error) {
	return me.Do("GET", "/users/"+id, nil)
}
