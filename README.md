```
$ http localhost:8080/proto.api.v1.APIService/Greet name=john
HTTP/1.1 200 OK
Accept-Encoding: gzip
Content-Encoding: gzip
Content-Length: 51
Content-Type: application/json
Date: Fri, 10 Feb 2023 15:07:53 GMT

{
    "greeting": "Hello, john!"
}
```

```
$ grpcurl -plaintext -d '{"name":"john"}' localhost:8080 proto.api.v1.APIService/Greet
{
  "greeting": "Hello, john!"
}
```
