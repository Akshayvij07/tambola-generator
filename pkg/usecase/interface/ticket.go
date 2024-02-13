package interfaces

import (
	"context"

	req "github.com/Akshayvij07/thambola-generator/pkg/api/handler/request"
	"github.com/Akshayvij07/thambola-generator/pkg/domain"
)

type TicketUseCase interface {
	GenerateSets(ctx context.Context, numSets int) ([]domain.TambolaSet, error)
	GetTicket(ctx context.Context, pagination req.Pagination) (Ticket []domain.Ticket, err error)
}
