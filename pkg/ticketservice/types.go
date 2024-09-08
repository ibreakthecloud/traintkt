package ticketservice

import (
	"sync"

	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
	"github.com/ibreakthecloud/traintkt/pkg/store"
)

type TicketServer struct {
	pb.UnimplementedTicketServiceServer
	mu    sync.Mutex
	Store store.Store
}
