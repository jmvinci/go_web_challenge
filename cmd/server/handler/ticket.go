package handler

import (
	"net/http"

	tickets "github.com/bootcamp-go/desafio-go-web/internal/ticket"
	"github.com/bootcamp-go/desafio-go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

/*
	type Service struct {
		service tickets.Service
	}

	func NewService(s tickets.Service) *Service {
		return &Service{
			service: s,
		}
	}
*/
type Ticket struct {
	sv *tickets.Service
}

func NewTicket(sv *tickets.Service) *Ticket {
	return &Ticket{sv: sv}
}

func (t *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		//tickets, err := s.service.GetTotalTickets(c, destination)
		tickets, err := (*t.sv).GetTotalTickets(c, destination)
		if err != nil {
			c.JSON(http.StatusOK, web.ErrorResponse{Message: "ok", Status: http.StatusInternalServerError, Code: err.Error()})
			return
		}

		c.JSON(http.StatusOK, web.Response{Message: "ok", Data: tickets})
	}
}

func (t *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := (*t.sv).AverageDestination(c, destination)
		if err != nil {
			c.JSON(http.StatusOK, web.ErrorResponse{Message: "ok", Status: http.StatusInternalServerError, Code: err.Error()})
			return
		}

		c.JSON(http.StatusOK, web.Response{Message: "ok", Data: avg})
	}
}
