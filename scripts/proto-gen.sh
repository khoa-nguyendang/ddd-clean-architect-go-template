protoc -I $1 --go_out=$1 --go_opt=paths=source_relative --go-grpc_out=$1 --go-grpc_opt=paths=source_relative $1/models/*.proto $1/services/*.proto

protoc -I src/internal/protos/v1 --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative src/internal/protos/v1/models/*.proto src/internal/protos/v1/services/*.proto 