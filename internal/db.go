package internal

import (
	"fmt"

	"github.com/sudip34/go-grpc-gateway-example/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

// NewDB creates a new array to mimic the behaviour of a in-memory database
func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

// AddOrder adds a new order to the DB collection. Returns an error on duplicate ids
func (mydb *DB) AddOrder(order *orders.Order) error {
	for _, myorder := range mydb.collection {
		if myorder.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id: %d", order.GetOrderId())
		}
	}
	mydb.collection = append(mydb.collection, order)
	return nil
}
