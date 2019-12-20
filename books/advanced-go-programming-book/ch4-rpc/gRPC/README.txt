4.4 gRPC入门
gRPC是Google公司基于Protobuf开发的跨语言的开源RPC框架。gRPC基于HTTP/2协议设计，可以基于一个HTTP/2链接提供多个服务，
对于移动设备更加友好。本节将讲述gRPC的简单用法。

4.5 gRPC进阶
作为一个基础的RPC框架，安全和扩展是经常遇到的问题。本节将简单介绍如何对gRPC进行安全认证。然后介绍通过gRPC的截取器特性，
以及如何通过截取器优雅地实现Token认证、调用跟踪以及Panic捕获等特性。最后介绍了gRPC服务如何和其他Web服务共存。

4.6 gRPC和Protobuf扩展
目前开源社区已经围绕Protobuf和gRPC开发出众多扩展，形成了庞大的生态。本节我们将简单介绍验证器和REST接口扩展。

protoc --go_out=plugins=grpc:. hello.proto