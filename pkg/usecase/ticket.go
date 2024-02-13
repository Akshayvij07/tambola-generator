package usecase

import (
	"context"
	"fmt"
	"sort"

	req "github.com/Akshayvij07/thambola-generator/pkg/api/handler/request"
	"github.com/Akshayvij07/thambola-generator/pkg/domain"
	interfaces "github.com/Akshayvij07/thambola-generator/pkg/repository/interface"
	services "github.com/Akshayvij07/thambola-generator/pkg/usecase/interface"
	"github.com/Akshayvij07/thambola-generator/pkg/utils"
)

type ticketUseCase struct {
	ticketRepo interfaces.TicketRepository
}

func NewTicketUseCase(repo interfaces.TicketRepository) services.TicketUseCase {
	return &ticketUseCase{
		ticketRepo: repo,
	}
}

// func (uc *ticketUseCase) GenerateSets(ctx context.Context, numSets int) ([]domain.Ticket, error) {
// 	var tickets []domain.Ticket

// 	// Generate N sets of Tambola tickets
// 	for i := 0; i < numSets; i++ {
// 		// Generate a single ticket
// 		ticket, err := generateSingleTicket()
// 		if err != nil {
// 			return nil, fmt.Errorf("error generating ticket: %v", err)
// 		}

// 		// Save the generated ticket to the database
// 		err = uc.ticketRepo.SaveTicket(ctx,&ticket)
// 		if err != nil {
// 			return nil, fmt.Errorf("error saving ticket to the database: %v", err)
// 		}

// 		// Add the generated ticket to the array
// 		tickets = append(tickets, ticket)
// 	}

// 	return tickets, nil
// }

func (uc *ticketUseCase) GenerateSets(ctx context.Context, numSets int) ([]domain.TambolaSet, error) {
	var tambolaSets []domain.TambolaSet
	valueMap := make(map[int]bool)
	checkMap := make(map[int][]int)
	// Generate N sets of Tambola tickets
	for i := 0; i < numSets; i++ {
		// Generate 6 tickets for each set
		var tickets []domain.Ticket
		for j := 0; j < 6; j++ {

			// Generate a single ticket
			ticket := utils.HandleDuplicates(utils.GenerateTicket(), valueMap, checkMap)
			fmt.Println(ticket)
			strTicket, err := utils.MatrixToString(ticket)

			if err != nil {
				return nil, fmt.Errorf("error generating ticket: %v", err)
			}

			//setId := utils.GenerateUniqueID(numSets)

			// // Save the generated ticket to the database
			// err = uc.ticketRepo.SaveTicket(ctx, ticket)
			// if err != nil {
			// 	return nil, fmt.Errorf("error saving ticket to the database: %v", err)
			// }

			// Add the generated ticket to the array
			tickets = append(tickets, domain.Ticket{Entries: strTicket})
		}

		//Sort the Ticket
		sort.Slice(tickets, func(i, j int) bool {
			return tickets[i].ID < tickets[j].ID
		})

		// Create a TambolaSet with the generated tickets
		tambolaSet := domain.TambolaSet{
			Tickets: tickets,
		}

		// Save the TambolaSet to the database
		err := uc.ticketRepo.SaveTambolaSet(ctx, &tambolaSet)
		if err != nil {
			return nil, fmt.Errorf("error saving TambolaSet to the database: %v", err)
		}

		// Add the generated TambolaSet to the array
		tambolaSets = append(tambolaSets, tambolaSet)
	}

	// for val := range tambolaSets{
	// 	uc.ticketRepo.SaveTicket(ctx,utils.MatrixToString(val))
	// }

	return tambolaSets, nil
}

// generateSingleTicket is a placeholder function for generating a single ticket
// func generateSingleTicket() (string, error) {

// 	// This is a placeholder and should be replaced with your actual logic

// 	// For demonstration purposes, a ticket with all zeros is generated
// 	// You should replace this with your actual logic to generate tickets
// 	valueMap := make(map[int]bool)
// 	ticket := utils.HandleDuplicates(utils.GenerateTicket(), valueMap)

// 	//utils.SortRows(ticket.Entries)
// 	fmt.Println(ticket)
// 	strTicket, err := utils.MatrixToString(ticket)

// 	// if utils.ValidateTambolaTicket(ticket.Entries) != nil {
// 	// 	return domain.Ticket{}, errors.New("generated ticket is not valid")
// 	// }

// 	// Assume the ticket entries are initialized to all zeros
// 	// Modify this based on your actual ticket generation logic

// 	return strTicket, err
// }

func (tc *ticketUseCase) GetTicket(ctx context.Context, pagination req.Pagination) (Ticket []domain.Ticket, err error) {
	Ticket, err = tc.ticketRepo.GetTicket(ctx, pagination)

	return Ticket, err
}
