package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type storage struct {
	Path string
}

type Storage interface {
	Get() ([]domain.Ticket, error)
	Set([]domain.Ticket) error
}

func NewStorage(path string) Storage {
	return &storage{Path: path}
}

func (s *storage) Get() (tickets []domain.Ticket, err error) {
	file, err := os.Open(s.Path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	for _, v := range data {
		newTicket, err := s.ticketFactory(v)
		if err != nil {
			break
		}
		tickets = append(tickets, newTicket)
	}
	return
}
func (s *storage) Set([]domain.Ticket) (err error) {
	return
}

func (s *storage) ticketFactory(ticket []string) (newTicket domain.Ticket, err error) {
	price, err := strconv.ParseFloat(ticket[5], 64)

	if err != nil {
		return
	}
	newTicket = domain.Ticket{
		Id:      ticket[0],
		Name:    ticket[1],
		Email:   ticket[2],
		Country: ticket[3],
		Time:    ticket[4],
		Price:   price,
	}
	return
}
