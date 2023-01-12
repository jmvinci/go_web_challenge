package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	middlewares "github.com/bootcamp-go/desafio-go-web/cmd/server/middleware"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	tickets "github.com/bootcamp-go/desafio-go-web/internal/ticket"
	"github.com/gin-gonic/gin"
)

type Router struct {
	db []domain.Ticket

	en *gin.Engine
}

func NewRouter(en *gin.Engine, storage *[]domain.Ticket) *Router {
	return &Router{en: en, db: *storage}
}

func (r *Router) SetRoutes() {
	r.MapRoutes()
}

func (r *Router) MapRoutes() {
	rp := tickets.NewRepository(r.db)
	sv := tickets.NewService(&rp)
	hd := handler.NewTicket(&sv)

	r.en.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	t := r.en.Group("ticket")
	t.Use(middlewares.AuthMiddleware())
	t.GET("getByCountry/:dest", hd.GetTicketsByCountry())
	t.GET("getAverage/:dest", hd.AverageDestination())

}
