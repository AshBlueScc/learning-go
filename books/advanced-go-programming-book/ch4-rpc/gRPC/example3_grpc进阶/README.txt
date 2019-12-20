4.5 gRPC进阶
作为一个基础的RPC框架，安全和扩展是经常遇到的问题。本节将简单介绍如何对gRPC进行安全认证。然后介绍通过gRPC的截取器特性，
以及如何通过截取器优雅地实现Token认证、调用跟踪以及Panic捕获等特性。最后介绍了gRPC服务如何和其他Web服务共存。

4.5.1 证书认证(证书的认证是针对每个gRPC链接的认证)
gRPC建立在HTTP/2协议之上，对TLS提供了很好的支持。我们前面章节中gRPC的服务都没有提供证书支持，因此客户端在链接服务器中通过grpc.
WithInsecure()选项跳过了对服务器证书的验证。没有启用证书的gRPC服务在和客户端进行的是明文通讯，信息面临被任何第三方监听的风险。
为了保障gRPC通信不被第三方监听篡改或伪造，我们可以对服务器启动TLS加密特性。

openssl下载地址: https://oomake.com/download/openssl

可以用以下命令为服务器和客户端分别生成私钥和证书：

genrsa -out server.key 2048
req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" -key server.key -out server.crt

genrsa -out client.key 2048
req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" -key client.key -out client.crt

.key为后缀名的是私钥文件，需要妥善保管。
.crt为后缀名的是证书文件，可以简单理解为公钥文件，并不需要秘密保存。
subj参数中的/CN=server.grpc.io表示服务器的名字为server.grpc.io,在验证服务器的证书时需要用到该信息

根证书的生成方式和自签名证书的生成方式类似：
genrsa -out ca.key 2048
req -new -x509 -days 3650 -subj "/C=GB/L=China/O=gobook/CN=github.com" -key ca.key -out ca.crt

重新对服务器端证书进行签名：
req -new -subj "/C=GB/L=China/O=server/CN=server.io" -key server.key -out server.csr
x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.crt

签名的过程中引入了一个新的以.csr为后缀名的文件，它表示证书签名请求文件。在证书签名完成之后可以删除.csr文件。
如果客户端的证书也采用CA根证书签名的话，服务器端也可以对客户端进行证书认证。我们用CA根证书对客户端证书签名：
req -new -subj "/C=GB/L=China/O=client/CN=client.io" -key client.key -out client.csr
x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.crt

将proto文件编译为:go文件，grpc库
protoc --go_out=plugins=grpc:. hello.proto

4.5.2 Token认证
gRPC还为每个gRPC方法调用提供了认证支持，这样就基于用户Token对不同的方法访问进行权限管理。

4.5.3 拦截器
gRPC中的grpc.UnaryInterceptor和grpc.StreamInterceptor分别对普通方法和流方法提供了拦截器的支持。我们这里简单介绍普通方法的拦截器用法。

4.5.4 和Web服务共存
gRPC构建在HTTP/2协议之上，因此我们可以将gRPC服务和普通的Web服务架设在同一个端口之上。

TLS(Transport Layer Security):传输层安全性协议