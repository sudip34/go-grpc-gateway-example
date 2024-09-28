package main

import (
	"log"
	"net"


  "github.com/sudip34/go-grpc-gateway-example/internal"
  "github.com/sudip34/go-grpc-gateway-example/protogen/golang/orders"
	"google.golang.org/grpc"
)


func main () {
	const addr = "0.0.0.0:50051"

	// from net create a TCP listner on the specified poet
	
	listner, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

  // create a gRPC server instance
	


	server := grpc.NewServer()

	// create an order service instance with a reference to the DB
	
	db := internal.NewDB()
	orderService := internal.NewOrderService(db)

	// register the order service with the grpc server
	
	orders.RegisterOrdersServer(server, &orderService)

	// start listening to requests
	
	log.Printf("Server listening at %v", listener.Addr())

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
