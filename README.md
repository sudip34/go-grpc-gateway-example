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

