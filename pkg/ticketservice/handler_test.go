package ticketservice

import (
	"context"
	"testing"

	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
	"github.com/ibreakthecloud/traintkt/pkg/store/inmemory"

	"github.com/stretchr/testify/assert"
)

func TestPurchaseTicket(t *testing.T) {
	store := inmemory.New()
	s := NewTicketServer(store)
	req := &pb.PurchaseTicketRequest{
		From:  "London",
		To:    "France",
		User:  &pb.User{FirstName: "Harsh", LastName: "Karn", Email: "harsh@karn.io"},
		Price: 20.0,
	}

	receipt, err := s.PurchaseTicket(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, req.From, receipt.From)
	assert.Equal(t, req.To, receipt.To)
	assert.Equal(t, req.User, receipt.User)
	assert.Equal(t, req.Price, receipt.Price)
	assert.NotEmpty(t, receipt.Seat)
}

func TestGetReceipt(t *testing.T) {
	store := inmemory.New()
	s := NewTicketServer(store)
	req := &pb.PurchaseTicketRequest{
		From:  "London",
		To:    "France",
		User:  &pb.User{FirstName: "Harsh", LastName: "Karn", Email: "harsh@karn.io"},
		Price: 20.0,
	}

	purchasedReceipt, _ := s.PurchaseTicket(context.Background(), req)

	getReceiptReq := &pb.GetReceiptRequest{Id: purchasedReceipt.Id}
	receipt, err := s.GetReceipt(context.Background(), getReceiptReq)
	assert.NoError(t, err)
	assert.Equal(t, purchasedReceipt, receipt)
}

func TestViewSeatAllocation(t *testing.T) {
	store := inmemory.New()
	s := NewTicketServer(store)
	req := &pb.PurchaseTicketRequest{
		From:  "London",
		To:    "France",
		User:  &pb.User{FirstName: "Harsh", LastName: "Karn", Email: "harsh@karn.io"},
		Price: 20.0,
	}

	s.PurchaseTicket(context.Background(), req)

	viewReq := &pb.ViewSeatAllocationRequest{Section: "A"}
	allocation, err := s.ViewSeatAllocation(context.Background(), viewReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, allocation.Allocations)
}

func TestRemoveUser(t *testing.T) {
	store := inmemory.New()
	s := NewTicketServer(store)
	req := &pb.PurchaseTicketRequest{
		From:  "London",
		To:    "France",
		User:  &pb.User{FirstName: "Harsh", LastName: "Karn", Email: "harsh@karn.io"},
		Price: 20.0,
	}

	receipt, _ := s.PurchaseTicket(context.Background(), req)

	removeReq := &pb.RemoveUserRequest{Id: receipt.Id}
	response, err := s.RemoveUser(context.Background(), removeReq)
	assert.NoError(t, err)
	assert.True(t, response.Success)

	// Verify that the user has been removed
	_, err = s.GetReceipt(context.Background(), &pb.GetReceiptRequest{Id: receipt.Id})
	assert.Error(t, err)
}

func TestModifySeat(t *testing.T) {
	store := inmemory.New()
	s := NewTicketServer(store)
	req := &pb.PurchaseTicketRequest{
		From:  "London",
		To:    "France",
		User:  &pb.User{FirstName: "Harsh", LastName: "Karn", Email: "harsh@karn.io"},
		Price: 20.0,
	}

	receipt, _ := s.PurchaseTicket(context.Background(), req)

	modifyReq := &pb.ModifySeatRequest{Id: receipt.Id, NewSeat: "B10"}
	response, err := s.ModifySeat(context.Background(), modifyReq)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "B10", response.NewSeat)

	// Verify that the seat has been modified
	updatedReceipt, _ := s.GetReceipt(context.Background(), &pb.GetReceiptRequest{Id: receipt.Id})
	assert.Equal(t, "B10", updatedReceipt.Seat)
}
