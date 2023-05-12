package gitlab

import "net/http"

const projectsPath = "/projects"

func (me *GitLab) GetProjects() (*http.Response, error) {
	return me.Do("GET", projectsPath, nil)
}

func (me *GitLab) GetProject(id string) (*http.Response, error) {
	return me.Do("GET", projectsPath+"/"+id, nil)
}
