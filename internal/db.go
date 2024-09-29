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

func (mydb *DB) GetOrderByID(orderID uint64) *orders.Order {
	
	for _, myorder := range mydb.collection {
		if myorder.OrderId == orderID {
			return myorder
		}
	}
	return nil
}


func (mydb *DB) GetOrdersByIds(orderIDs []uint64) []*orders.Order {
	filtered := make([]*orders.Order, 0)
	for _, idx := range orderIDs {
   
		for _, myorder := range mydb.collection {
			if myorder.OrderId == idx {
				filtered = append(filtered,myorder)
				break
			}
		}
	}
	return filtered 
}


func (d *DB) UpdateOrder(order *orders.Order) {
	for i, o := range d.collection {
		if o.OrderId == order.OrderId {
			d.collection[i] = order
			return
		}
	}
}

func (mydb *DB) RemoveOrder(orderID uint64) {
	filtered := make([]*orders.Order, 0, len(mydb.collection)-1)
   
	for i := range mydb.collection {
		if mydb.collection[i].OrderId != orderID {
			filtered = append(filtered,mydb.collection[i])
	 }
 }
	mydb.collection = filtered
}
