.PHONY: gen-protos
gen-protos:
	@for protoDir in `find src/internal/protos/* -type f -regex ".*/models/.*\.proto" -exec dirname {} \; | sort | uniq`; do \
		echo "protoc for dir $${protoDir}" ; \
		protoc --go_out=paths=source_relative:./src/core/models/transfers --proto_path=./src/internal/protos $${protoDir}/*.proto; \
	done

	@for protoDir in `find src/internal/protos/* -type f -regex ".*/services/.*\.proto" -exec dirname {} \; | sort | uniq`; do \
		echo "protoc for dir $${protoDir}" ; \
		protoc --go_out=paths=source_relative:./src/infrastructure/microservices_accessors --proto_path=./src/internal/protos $${protoDir}/*.proto; \
	done
