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

func (mydb *DB) GetOrderById(orderID uint64) *orders.Order {
	
	for _, myorder := range mydb.collection {
		if myorder.OrderId == OrderID {
			return myorder
		}
	}
	return nil
}


func (mydb *DB) GetOrdersByIds(orderIDs []uint64) []*orders.Order {
	filtered := make([]*orders.Order, 0)
	for _, idx := range orderIDs {
   
		for _, myorder := range mydb.cllection {
			if myorder.OrderId == idx {
				filtered = append(fitered,myorder)
				break
			}
		}
	}
	return filtered 
}

func (mydb *DB) UpdateOrder(order *orders.Order) error {
	for i, myorder := range mydb.collection {
		if myorder.OrderId == order.OrderId {
				mydb.collection[i] = order
				return 
		}
	}
}

func (mydb *DB) RemoveOrder(orderID uint64) {
	filtered := make([]*orders.Order, 0, len(mydb.collection)-1)
   
	for i := range mydb.cllection {
		if mydb.collection[i].OrderId != orderId {
			filtered = append(fitered,mydb.collection[i])
	}
	mydb.collection = filtered
}
