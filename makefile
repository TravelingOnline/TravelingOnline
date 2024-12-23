gen-proto:
		protoc -I=./auth/protobufs --go_out=./auth/protobufs --go_opt=paths=source_relative  --go-grpc_out=./auth/protobufs --go-grpc_opt=paths=source_relative ./auth/protobufs/*.proto
		protoc -I=./hotel/protobufs --go_out=./hotel/protobufs --go_opt=paths=source_relative  --go-grpc_out=./hotel/protobufs --go-grpc_opt=paths=source_relative ./hotel/protobufs/*.proto
		protoc -I=./bank/protobufs --go_out=./bank/protobufs --go_opt=paths=source_relative  --go-grpc_out=./bank/protobufs --go-grpc_opt=paths=source_relative ./bank/protobufs/*.proto