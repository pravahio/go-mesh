JS RPC
```
grpc_tools_node_protoc -I=./../ --js_out=import_style=commonjs,binary:. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` rpc.proto
```

GO RPC
```
protoc rpc.proto --go_out=plugins=grpc:.
```