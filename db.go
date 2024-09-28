package internal

import(
	"fmt"

  "github.com/sudip34/go-grpc-gateway-example/protogen/golang/orders"
)


type DB struct {
	collection []*orders.order 
}

func NewDB() *DB {
	return &DB {
		collection: make([]*orders.Order, 0),
	}
}

func (mydb *DB) AddOrder(order *orders.Order) error {
	for _, myorder := range mydb.collection {
		if myorder.OrderId == order.OrderId {
			return fmt.Errorf("duplicate Order id: %d", order.GetOrderId())
		}
	}

  mydb.collection = append(mydb.collection, order)
	return nil
}
