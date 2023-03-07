protoc -I internal/protos/gateway/ -I internal/includes/googleapis/ -I internal/includes/protoc-gen-swagger/ --go_out="internal/protos/gateway/" --go_opt=paths=source_relative --go-grpc_out=internal/protos/gateway/ --go-grpc_opt=paths=source_relative internal/protos/gateway/combine.proto & protoc -I internal/protos/gateway/ -I internal/includes/googleapis/ -I internal/includes/protoc-gen-swagger/ --grpc-gateway_out internal/protos/gateway/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true internal/protos/gateway/combine.proto  & protoc -I internal/protos/gateway/ -I internal/includes/googleapis/ -I internal/includes/protoc-gen-swagger/ --openapiv2_out internal/protos/gateway/ --openapiv2_opt logtostderr=true internal/protos/gateway/combine.proto