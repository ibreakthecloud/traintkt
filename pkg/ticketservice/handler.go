package ticketservice

import (
	"context"
	"fmt"

	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *TicketServer) PurchaseTicket(ctx context.Context, req *pb.PurchaseTicketRequest) (*pb.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%s-%s-%s", req.From, req.To, req.User.Email)
	seat := s.Store.AllocateSeat(MAX_SEATS_PER_SECTION)
	if seat == "" {
		return nil, status.Errorf(codes.NotFound, "No available seats")
	}

	receipt := &pb.Receipt{
		Id:    id,
		From:  req.From,
		To:    req.To,
		User:  req.User,
		Price: req.Price,
		Seat:  seat,
	}

	s.Store.AddReceipt(receipt)
	s.Store.ReserveSeat(seat, req.User)

	return receipt, nil
}

func (s *TicketServer) GetReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt := s.Store.GetReceipt(req.Id)
	if receipt == nil {
		return nil, status.Errorf(codes.NotFound, "Receipt not found")
	}

	return receipt, nil
}

func (s *TicketServer) ViewSeatAllocation(ctx context.Context, req *pb.ViewSeatAllocationRequest) (*pb.SeatAllocationResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Store.GetSeatAllocation(req.Section)
}

func (s *TicketServer) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt := s.Store.GetReceipt(req.Id)
	if receipt == nil {
		return &pb.RemoveUserResponse{Success: false}, status.Errorf(codes.NotFound, "User not found")
	}

	s.Store.DeleteSeat(receipt.Seat)
	s.Store.DeleteReceipt(req.Id)

	return &pb.RemoveUserResponse{Success: true}, nil
}

func (s *TicketServer) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt := s.Store.GetReceipt(req.Id)
	if receipt == nil {
		return &pb.ModifySeatResponse{Success: false}, status.Errorf(codes.NotFound, "User not found")
	}

	if !s.Store.IsSeatAvailable(req.NewSeat) {
		return &pb.ModifySeatResponse{Success: false}, status.Errorf(codes.AlreadyExists, "Seat already occupied")
	}

	s.Store.DeleteSeat(receipt.Seat)

	s.Store.ReserveSeat(req.NewSeat, receipt.User)
	s.Store.UpdateReceiptWithSeat(receipt.Id, req.NewSeat)

	return &pb.ModifySeatResponse{Success: true, NewSeat: req.NewSeat}, nil
}
