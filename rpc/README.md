JS RPC
```
grpc_tools_node_protoc -I=./../ --js_out=import_style=commonjs,binary:. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` rpc.proto
```

```
protoc -I=./../ rpc.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
```

GO RPC
```
protoc rpc.proto --go_out=plugins=grpc:.
```