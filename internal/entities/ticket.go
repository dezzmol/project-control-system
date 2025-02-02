package entities

type Ticket struct {
	Id          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Status      string `db:"status"`
}

type TicketDTO struct {
	TicketId    string `json:"ticketId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UserTicket struct {
	TicketId string `json:"ticketId"`
	Title    string `json:"title"`
	Status   string `json:"status"`
}

type CreateTicket struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AssignTicketDTO struct {
	UserId string `json:"userId"`
}

type TicketStatusDTO struct {
	Status string `json:"status"`
}
