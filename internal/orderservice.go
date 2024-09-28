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
	orders.UnimplementedOrdersServer
}

// NewOrderServie creates a new OrderService

func NewOrderService(mydb *DB) OrderService {
	return OrderService {db: mydb}
}

// AddOrder
// we will implement the AddOrder method of grpc OrderServer interface to add a new orders

func (or *OrderService) AddOrder(_ context.Context, request *orders.PayloadWithSingleOrder) (*orders.Empty, error){
		log.Printf("Recevied an add-order request")

		err := or.db.AddOrder(request.GetOrder())

		return &orders.Empty{}, err
}

func (or *OrderService) GetOrder(_ context.Context, request *orders.PayloadWithOrderID) (*orders.PayloadWithSingleOrder, error){
		log.Printf("Recevied a Get-order request")

		order := or.db.GetOrderById(req.GetOrderId())

		if order == nil {
			return nil, fmt.Errorf("order not found for orderID: %d", req.GetOrderId())
		}

		return &order.PayloadWithSingleOrder{Order: order}, nil
}


func (or *OrderService) UpdateOrder(_ context.Context, request *orders.PayloadWithSingleOrder) (*orders.Empty, error){
		log.Printf("Recevied a Update-order request")

		return &order.Empty{},nil
}


// RemoveOrder implements the RemoveOrder method of the grpc OrdersServer interface to remove an order
func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.Empty, error) {
	log.Printf("Received a remove order request")
	o.db.RemoveOrder(req.GetOrderId())
	return &orders.Empty{}, nil
}
