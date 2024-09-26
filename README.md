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
