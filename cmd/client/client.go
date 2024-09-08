package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Purchase a ticket
	r, err := c.PurchaseTicket(ctx, &pb.PurchaseTicketRequest{
		From: "London",
		To:   "France",
		User: &pb.User{
			FirstName: "Harsh",
			LastName:  "Karn",
			Email:     "harsh@karn.io",
		},
		Price: 20.0,
	})
	if err != nil {
		log.Fatalf("could not purchase ticket: %v", err)
	}
	log.Printf("Ticket purchased: %s", r.Id)

	// Get receipt
	receipt, err := c.GetReceipt(ctx, &pb.GetReceiptRequest{Id: r.Id})
	if err != nil {
		log.Fatalf("could not get receipt: %v", err)
	}
	log.Printf("Receipt: %v", receipt)

	// View seat allocation
	seatAllocation, err := c.ViewSeatAllocation(ctx, &pb.ViewSeatAllocationRequest{Section: "A"})
	if err != nil {
		log.Fatalf("could not view seat allocation: %v", err)
	}
	log.Printf("Seat allocation for section A: %v", seatAllocation)

	// Modify seat
	modifySeatResponse, err := c.ModifySeat(ctx, &pb.ModifySeatRequest{Id: r.Id, NewSeat: "B10"})
	if err != nil {
		log.Fatalf("could not modify seat: %v", err)
	}
	log.Printf("Seat modified: %v", modifySeatResponse)

	// Remove user
	removeUserResponse, err := c.RemoveUser(ctx, &pb.RemoveUserRequest{Id: r.Id})
	if err != nil {
		log.Fatalf("could not remove user: %v", err)
	}
	log.Printf("User removed: %v", removeUserResponse)
}
