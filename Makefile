install-pb-deps:
				go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
				go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-pb:
				protoc --proto_path=./proto --go_out=paths=source_relative:./gen $(find ./proto -name "*.proto")
