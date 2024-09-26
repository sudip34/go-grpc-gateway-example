package main

import (
	"fmt"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
	
	"github.com/sudip34/go-grpc-gateway-example/protogen/golang/product"
	"github.com/sudip34/go-grpc-gateway-example/protogen/golang/orders"

)

func main() {

	orderItem := orders.Order {
    OrderId: 10,
		CustomerId: 11,
		isActive: true,
		OrderDate: &date.Date{year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "Benana", ProductType: product.ProductType_FOOD},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)

	if err != nil {
		log.Fatal("deserialization error: ", err)
	}

	fmt.Println(string(bytes))

}
