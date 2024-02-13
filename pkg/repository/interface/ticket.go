package interfaces

import (
	"context"

	req "github.com/Akshayvij07/thambola-generator/pkg/api/handler/request"
	"github.com/Akshayvij07/thambola-generator/pkg/domain"
)

type TicketRepository interface {
	SaveTicket(ctx context.Context, ticket string) error
	SaveTambolaSet(ctx context.Context, Set *domain.TambolaSet) error
	GetTicket(ctx context.Context, pagination req.Pagination) (Ticket []domain.Ticket, err error)
}
