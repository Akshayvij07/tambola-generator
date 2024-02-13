package handler

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"

	req "github.com/Akshayvij07/thambola-generator/pkg/api/handler/request"
	res "github.com/Akshayvij07/thambola-generator/pkg/api/handler/respondse"
	services "github.com/Akshayvij07/thambola-generator/pkg/usecase/interface"
)

type TicketHandler struct {
	ticketUseCase services.TicketUseCase
}

func NewUserHandler(usecase services.TicketUseCase) *TicketHandler {
	return &TicketHandler{
		ticketUseCase: usecase,
	}
}

func (cr *TicketHandler) CreateTicket(ctx *gin.Context) {
	//paramsId := ctx.Param("id")
	// id, err := strconv.Atoi(paramsId)
	// if err != nil {
	// 	res.ErrrorResponse(ctx, 402, "id params not found", err.Error())
	// 	return
	// }
	var ticket req.NumberOfTicket
	if err := ctx.Bind(&ticket); err != nil {
		res.ErrrorResponse(ctx, 402, "failed to bind", err.Error())
		return
	}
	Token, err := cr.ticketUseCase.GenerateSets(ctx, ticket.Count)
	fmt.Println(Token)
	if err != nil {
		res.ErrrorResponse(ctx, 400, "failed to create Token", err.Error())
		return
	}
	res.SuccessResponse(ctx, 200, "Token Number", Token)

}

func (cr *TicketHandler) AllTicket(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {

		page = 1
	}

	perPage, err := strconv.Atoi(ctx.Query("perPage"))
	if err != nil {

		perPage = 10
	}

	ListofOrders := req.Pagination{
		Page:    uint(page),
		PerPage: uint(perPage),
	}

	tickets, err := cr.ticketUseCase.GetTicket(ctx, ListofOrders)
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i].ID < tickets[j].ID
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't List the Tickets",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Tickets ",
		Data:       tickets,
		Errors:     nil,
	})
}
