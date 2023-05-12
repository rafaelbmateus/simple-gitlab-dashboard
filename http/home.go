package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (me *Handler) Home(c *gin.Context) {
	projects, err := me.Usecase.GetProjects()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	users, err := me.Usecase.GetUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title":    "Home",
		"host":     me.Usecase.GitLab.Host,
		"projects": projects,
		"users":    users,
	})
}
