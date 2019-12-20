grpcurl是Go语言开源社区开发的工具，需要手工安装：

go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl

grpcurl -plaintext localhost:1234 list
grpcurl -plaintext localhost:1234 list proto.HelloService
grpcurl -plaintext localhost:1234 describe proto.HelloService
grpcurl -plaintext localhost:1234 describe proto.String

grpcurl -plaintext -d "{\"value\":\"World!\"}" localhost:1234 proto.HelloService/Hello