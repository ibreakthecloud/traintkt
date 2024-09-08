package inmemory

import (
	"fmt"

	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
)

func New() *InMemory {
	return &InMemory{
		Booking: Booking{
			Receipts: make(map[string]*pb.Receipt),
			Seats:    make(map[string]*pb.User),
		},
	}
}

func (i *InMemory) AddReceipt(receipt *pb.Receipt) {
	i.Booking.Receipts[receipt.Id] = receipt
}

func (i *InMemory) ReserveSeat(seat string, user *pb.User) {
	i.Booking.Seats[seat] = user
}

func (i *InMemory) GetReceipt(id string) *pb.Receipt {
	return i.Booking.Receipts[id]
}

func (i *InMemory) GetSeatAllocation(section string) (*pb.SeatAllocationResponse, error) {
	var allocations []*pb.SeatAllocation

	for seat, user := range i.Booking.Seats {
		if (section == "A" && seat[0] == 'A') || (section == "B" && seat[0] == 'B') {
			allocations = append(allocations, &pb.SeatAllocation{
				Seat: seat,
				User: user,
			})
		}
	}

	return &pb.SeatAllocationResponse{Allocations: allocations}, nil
}

func (i *InMemory) IsSeatAvailable(seat string) bool {
	_, occupied := i.Booking.Seats[seat]
	return !occupied
}

func (i *InMemory) UpdateReceiptWithSeat(receiptId, newSeat string) {
	receipt := i.Booking.Receipts[receiptId]
	receipt.Seat = newSeat
	// ideally this is not required is we are updating the reference
	// todo: check if this is required and remove if not
	i.Booking.Receipts[receiptId] = receipt
}

func (i *InMemory) DeleteSeat(seat string) {
	delete(i.Booking.Seats, seat)
}

func (i *InMemory) DeleteReceipt(id string) {
	delete(i.Booking.Receipts, id)
}

func (i *InMemory) AllocateSeat(maxSeat int) string {
	sections := []string{"A", "B"}
	for _, section := range sections {
		for j := 1; j <= maxSeat; j++ {
			seat := fmt.Sprintf("%s%d", section, j)
			_, occupied := i.Booking.Seats[seat]
			if !occupied {
				return seat
			}
		}
	}
	return ""
}
