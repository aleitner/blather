protoc -I=../ --go_out=plugins=grpc:../internal/protobuf/ ../internal/protobuf/*.proto
