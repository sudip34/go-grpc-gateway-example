package main


import (
	"context"
  "fmt"
  "log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
  "github.com/sudip34/go-grpc-gateway-example/protogen/golang/orders";
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
 )

 func main() {
	 // Setup a connection to the server
	 orderServiceAddr := "localhost:50051"
	 conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	 if err != nil {
		 log.Fatalf("Connection with order service is not possible: %v", err)
	 }

	 // register gRPC server endpoint
	 // Note : First check gRPC server is runnign and accessiable

	 mux := runtime.NewServeMux()
	 if err = orders.RegisterOrdersHandler(context.Background(),mux, conn); err !=nil {
		 log.Fatalf("failed to register to the server %v", err)
	 }

	 // Listenig to request from the gateway server
	 addr := "0.0.0.0:8080"
	 fmt.Println("API gateway server is running on " + addr)
	 if err = http.ListenAndServe(addr, mux); err != nil {
		 log.Fatal("gateway server closed abruptly: ", err)
	 }
 }
