package gitlab

import "net/http"

func (me *GitLab) GetProjectCommits(id string) (*http.Response, error) {
	return me.Do("GET", "/projects/"+id+"/repository/commits", nil)
}
