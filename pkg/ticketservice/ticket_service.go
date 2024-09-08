package ticketservice

import (
	"github.com/ibreakthecloud/traintkt/pkg/store"
)

// assuming 50 seats per section
const MAX_SEATS_PER_SECTION = 50

func NewTicketServer(store store.Store) *TicketServer {
	return &TicketServer{
		Store: store,
	}
}
