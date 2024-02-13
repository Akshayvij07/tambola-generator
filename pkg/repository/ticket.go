package repository

import (
	"context"
	"fmt"

	req "github.com/Akshayvij07/thambola-generator/pkg/api/handler/request"
	"github.com/Akshayvij07/thambola-generator/pkg/domain"
	interfaces "github.com/Akshayvij07/thambola-generator/pkg/repository/interface"
	"gorm.io/gorm"
)

type ticketDatabase struct {
	DB *gorm.DB
}

func NewTicketRepository(DB *gorm.DB) interfaces.TicketRepository {
	return &ticketDatabase{DB}
}

// func (r *ticketDatabase) SaveTicket(ctx context.Context, ticket *domain.Ticket) error {
// 	// Prepare the SQL query for inserting a ticket
// 	query := `
// 	INSERT INTO tickets (
// 		entry_0_0, entry_0_1, entry_0_2, entry_0_3, entry_0_4, entry_0_5, entry_0_6, entry_0_7, entry_0_8,
// 		entry_1_0, entry_1_1, entry_1_2, entry_1_3, entry_1_4, entry_1_5, entry_1_6, entry_1_7, entry_1_8,
// 		entry_2_0, entry_2_1, entry_2_2, entry_2_3, entry_2_4, entry_2_5, entry_2_6, entry_2_7, entry_2_8
// 	) VALUES (
// 		$1, $2, $3, $4, $5, $6, $7, $8, $9,
// 		$10, $11, $12, $13, $14, $15, $16, $17, $18,
// 		$19, $20, $21, $22, $23, $24, $25, $26, $27
// 	) RETURNING id;
// 	`

// 	// Execute the SQL query with the values from the Ticket struct
// 	err := r.DB.Raw(
// 		query,
// 		ticket.Entries[0][0], ticket.Entries[0][1], ticket.Entries[0][2], ticket.Entries[0][3], ticket.Entries[0][4], ticket.Entries[0][5], ticket.Entries[0][6], ticket.Entries[0][7], ticket.Entries[0][8],
// 		ticket.Entries[1][0], ticket.Entries[1][1], ticket.Entries[1][2], ticket.Entries[1][3], ticket.Entries[1][4], ticket.Entries[1][5], ticket.Entries[1][6], ticket.Entries[1][7], ticket.Entries[1][8],
// 		ticket.Entries[2][0], ticket.Entries[2][1], ticket.Entries[2][2], ticket.Entries[2][3], ticket.Entries[2][4], ticket.Entries[2][5], ticket.Entries[2][6], ticket.Entries[2][7], ticket.Entries[2][8],
// 	).Scan(&ticket.ID)

// 	if err != nil {
// 		return fmt.Errorf("error inserting ticket into the database: %v", err)
// 	}

// 	return nil
// }

func (c *ticketDatabase) SaveTicket(ctx context.Context, ticket string) error {
	//newClinic := domain.Clinic{Name: data.Name, Location: data.Location}
	tk := domain.Ticket{Entries: ticket}
	err := c.DB.Create(&tk).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *ticketDatabase) SaveTambolaSet(ctx context.Context, Set *domain.TambolaSet) error {
	tambolaSet := domain.TambolaSet{Tickets: Set.Tickets}

	err := c.DB.Create(&tambolaSet).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *ticketDatabase) GetTicket(ctx context.Context, pagination req.Pagination) (Ticket []domain.Ticket, err error) {
	limit := pagination.PerPage
	offset := (pagination.Page - 1) * limit

	fmt.Println(limit, offset)
	query := `SELECT * FROM tickets ORDER BY Entries  DESC LIMIT $1 OFFSET $2`
	err = c.DB.Raw(query, limit, offset).Scan(&Ticket).Error
	return Ticket, err
}
