package port

import (
	"context"

	"github.com/iulianclita/json-ports/internal/port/domain"
)

type Repository interface {
	UpsertPort(ctx context.Context, port *domain.Port) error
	GetPorts(ctx context.Context) ([]*domain.Port, error)
}
