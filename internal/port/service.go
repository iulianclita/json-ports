package port

import (
	"context"
	"fmt"

	"github.com/iulianclita/json-ports/internal/port/domain"
)

type Service interface {
	UpsertPort(ctx context.Context, port *domain.Port) error
	GetPorts(ctx context.Context) ([]*domain.Port, error)
}

type portService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &portService{
		repo: repo,
	}
}

func (ps *portService) UpsertPort(ctx context.Context, port *domain.Port) error {
	if err := ps.repo.UpsertPort(ctx, port); err != nil {
		return fmt.Errorf("failed to upsert port with id = %s: %w", port.ID, err)
	}

	return nil
}

func (ps *portService) GetPorts(ctx context.Context) ([]*domain.Port, error) {
	ports, err := ps.repo.GetPorts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ports: %w", err)
	}

	return ports, nil
}
