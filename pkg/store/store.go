package store

import (
	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
)

type Store interface {
	AddReceipt(receipt *pb.Receipt)
	ReserveSeat(seat string, user *pb.User)
	GetSeatAllocation(section string) (*pb.SeatAllocationResponse, error)
	GetReceipt(id string) *pb.Receipt
	DeleteSeat(seat string)
	DeleteReceipt(id string)
	IsSeatAvailable(seat string) bool
	UpdateReceiptWithSeat(receiptId, newSeat string)
	AllocateSeat(maxSeat int) string
}
