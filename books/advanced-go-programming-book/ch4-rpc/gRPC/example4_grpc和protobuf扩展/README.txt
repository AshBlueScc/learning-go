验证器：
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
protoc -IC:/Go/bin/pkg/mod/ -IC:/Go/bin/src/github.com --proto_path=. --govalidators_out=. --go_out=plugins=grpc:. hello.proto


Rest:
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
protoc -I. -IC:/Go/bin/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.1/third_party/googleapis -IC:/Go/bin/src/github.com --grpc-gateway_out=. --go_out=plugins=grpc:. hello.proto

http://localhost:8080/get/gopher    返回：{"value":"Get: gopher"}
http://localhost:8080/post 参数：{"value":"grpc"}          返回：{ "value": "Post: grpc"}
Get请求可以在浏览器，post请求在postman
大致就是serService监听8080端口把请求转发到5000端口，grpcService监听5000端口处理请求并返回