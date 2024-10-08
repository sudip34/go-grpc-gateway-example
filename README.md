# go-grpc-gateway-example


## Description for PROTOCOL-BUFFER
| PROTOBUF | in GO |
| --- | ---|
|  `syntax`  | ***version*** used for the PROTOCOL-BUFFER|
|  `go_package` | specify the `improt path` for the generated GO language bindings. So, the complied code for `order.proto` will be a file path `prtogen/golang/orders/order.pb.go`|
|  `message` | structured unit that represent ***DATA***. Equivalent in GO is `struct` |
|  `uint64` | corresponds to `uint64` in GO |
|  `JSON`  | can ***optionally*** be specified, ensure _serialized messages_ align with defined field names. gRPC defaults to ***PascalCase*** |
|  `repeated` | Used for a ***Collection*** in general e.g Arrays or list.
|  `Enums` | enum |


To To compile this code we need to copy the `Date` definition file and add it to our project. You can create a folder called `google/api` under the `proto` folder and copy the code under the filename `date.proto`.

### Lets compile the code 

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
> this will set defalut

- $GOPATH/bin (if $GOPATH is set), or
- $HOME/go/bin (the default location for Go binaries).

> if not the path default updated then for bash 
```
ls $HOME/go/bin/protoc-gen-go
```

> add Go bin directory to `PATH` . open bash `nano ~/.bash` or other shell if your are using

```
export PATH=$PATH:$HOME/go/bin 

```
> Reload the shell

``` 
source ~/.bashrc 

```

> verify 

```
protoc-gen-go --version

```

***create a Makefile*** to compile the `proto` files
```
# Makefile

protoc:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	./**/*.proto
```


> `--go_out` is the output directory
> `--go_opt`  to specify Go package paths
> `./**/*.proto` glob expands the current directory and includes all `proto` files for the compilation process.

## Run makefile 

```
make protoc 
```


### for main.go()

- first install the `protobuf` & `genproto`

```
go get google.golang.org/protobuf 
go get google.golang.org/genproto 

```





### For now we use _service_ for method or function in PROTOCOL-BUFFER 

- `service`

- in a `service` we can use the object  as `message`

- To compile `service` definition we need a specific gRPC-specific binary `protoc-gen-go-grpc`


```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
``` 


### Install gRPC dependency


``` 
go get google.golang.org/grpc 
``` 


### update the Makefile

```
protoc:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
  --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
  ./**/*.proto
```


```
make protoc 

```



> create 2 directories `cmd/server` and `internal`

```
mkdir -p cmd/server  internal 
touch cmd/server/main.go internal/{db,orderservice}.go 
```



## Set up the API Gateway 

> gRPC does not supprt client with RESTful JSON architecture well
> to support we will use google library ***grpc-gateway*** [https://github.com/grpc-ecosystem/grpc-gateway]

```
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

```

- Now We need to add two new files into our `proto/google/api`  : `annotation.proto` , `http.proto`
- we need to make small adjustments to our service definition 
```
curl -L https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o proto/google/api/annotations.proto
curl -L https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o proto/google/api/http.proto 
```



#### Now will craete protobuf for the gRPC - Gateway and for that will make change in the Makefile

```
protoc:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./**/*.proto  
```


## Client Server 

#### Create API gateway server
```
mkdir cmd/client
touch cmd/client/main.go 

```

### Now we need to update 3 following files

- order.proto
- db.go 
- orderservice.go 

for the following request:

- GetOrder
- UpdateOrder
- RemoveOrder
