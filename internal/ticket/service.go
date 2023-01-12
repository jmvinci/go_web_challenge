package tickets

import (
	"context"
)

type service struct {
	rp Repository
}

type Service interface {
	GetTotalTickets(ctx context.Context, country string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

func NewService(rp *Repository) Service {
	return &service{rp: *rp}
}

func (s *service) GetTotalTickets(ctx context.Context, country string) (int, error) {
	tickets, err := s.rp.GetTicketByDestination(ctx, country)
	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.rp.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0.0, err
	}
	total, err := s.rp.GetAll(ctx)
	if err != nil {
		return 0.0, err
	}

	avg := float64(len(tickets)) / float64(len(total))
	return avg, nil
}
