# Train Ticket Booking System

This is a simple train ticket booking system implemented using Go and gRPC. The system allows users to purchase tickets, view receipts, allocate seats, and manage bookings.

## Features

- Purchase train tickets
- View ticket receipts
- View seat allocations by section
- Modify user seats
- Remove users from the train

## Storage Mechanism

The system uses an in-memory storage mechanism to store user and ticket information. The storage is implemented using a map data structure. Map because it provides constant time complexity, `O(1)`, for read, write, and delete operations.
System also have store layer as an interface,
```go
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
```
This interface is implemented by `InMemoryStore` which is used by the `TicketService` to store and retrieve data. 

To add a persistent storage mechanism, we can implement the `Store` interface using a database or file system.

## Prerequisites
- Go 
- Protocol Buffers compiler (`protoc`)
- Go plugins for the protocol compiler:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/ibreakthecloud/traintkt.git
   cd traintkt
   ```

2. Install the required Go packages:
   ```
   go mod tidy
   ```

3. Generate the gRPC code:
   ```
   protoc --go_out=pkg --go_opt=paths=source_relative \
          --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
          proto/ticket.proto
   ```

## Running the Application

### Start the server:

```
go run cmd/server/main.go
```

The server will start and listen on `localhost:50051`.

### Run the client:

```
go run cmd/client/main.go
```

The client will perform a series of operations demonstrating the functionality of the ticket booking system.

## Running Tests

To run the unit tests:

```bash
go test -v -coverprofile=cover.out ./... 
```
```bash
?       github.com/ibreakthecloud/traintkt/pkg/store    [no test files]
        github.com/ibreakthecloud/traintkt/cmd/client           coverage: 0.0% of statements
        github.com/ibreakthecloud/traintkt/pkg/store/inmemory           coverage: 0.0% of statements
        github.com/ibreakthecloud/traintkt/cmd/server           coverage: 0.0% of statements
        github.com/ibreakthecloud/traintkt/pkg/proto            coverage: 0.0% of statements
=== RUN   TestPurchaseTicket
--- PASS: TestPurchaseTicket (0.00s)
=== RUN   TestGetReceipt
--- PASS: TestGetReceipt (0.00s)
=== RUN   TestViewSeatAllocation
--- PASS: TestViewSeatAllocation (0.00s)
=== RUN   TestRemoveUser
--- PASS: TestRemoveUser (0.00s)
=== RUN   TestModifySeat
--- PASS: TestModifySeat (0.00s)
PASS
coverage: 89.7% of statements
ok      github.com/ibreakthecloud/traintkt/pkg/ticketservice    0.290s  coverage: 89.7% of statements
```
## Coverage

89.7% of statements are covered by the unit tests.