package http

import (
	"log"

	"github.com.br/rafaelbmateus/gitlab-dashboard/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router  *gin.Engine
	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	router := gin.Default()
	return &Handler{
		Router:  router,
		Usecase: usecase,
	}
}

func (me *Handler) Server() {
	me.Routes()

	if err := me.Router.Run(":3000"); err != nil {
		log.Fatalf("error on run server")
	}
}

func (me *Handler) Routes() {
	me.Router.LoadHTMLGlob("views/*")
	me.Router.Static("/assets", "./assets")
	me.Router.GET("/", me.Home)
}
