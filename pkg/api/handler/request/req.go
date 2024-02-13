package req

type NumberOfTicket struct {
	Count int
}

type Pagination struct {
	Page    uint `json:"page"`
	PerPage uint `json:"page_per"`
}
