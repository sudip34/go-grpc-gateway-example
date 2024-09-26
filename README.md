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

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

