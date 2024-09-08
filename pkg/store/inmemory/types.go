package inmemory

import (
	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
)

// Store to store metrics
type InMemory struct {
	Booking Booking
}

type Booking struct {
	Receipts map[string]*pb.Receipt
	Seats    map[string]*pb.User
}
