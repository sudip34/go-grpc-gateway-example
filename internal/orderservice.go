package internal

import (
	"context"
	"log"

  "github.com/sudip34/go-grpc-gateway-example/protogen/golang/orders"
)


// Here will implemented the grpc generated OrderServer interface  for our OrderService
// UnimplementedOrderServer must be embedded to have forwarded compatible implementation
//

type OrderService struct {
	db *DB
	orders.UnimplementedOrderServer
}

// NewOrderServie creates a new OrderService

func NewOrderServie(mydb *DB) OrderServie {
	return OrderService {db: mydb}
}

// AddOrder
// we will implement the AddOrder method of grpc OrderServer interface to add a new orders

func (or *OrderService) AddOrder(_ context.Context, request *orders.PayloadWithSingleOrder) (*orders.Empty, error){
		log.Printf("Recevied an add-order request")

		err := or.db.AddOrder(request.GetOrder())

		return &orders.Empty{}, err
}
