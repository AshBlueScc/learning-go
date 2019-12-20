梳理一下过程：
key私钥（生成） -> crt(证书，类似于公钥)

ca.key         ->       ca.crt


server.key  ->      server.csr  (中间文件)

ca.crt  +   ca.key  ->(把server.csr)     输出为 server.crt

同理可得    client.crt