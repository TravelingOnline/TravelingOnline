gen-proto:
		protoc -I=./auth/protobufs --go_out=./auth/protobufs --go_opt=paths=source_relative  --go-grpc_out=./auth/protobufs --go-grpc_opt=paths=source_relative ./auth/protobufs/*.proto
		protoc -I=./hotel/protobufs --go_out=./hotel/protobufs --go_opt=paths=source_relative  --go-grpc_out=./hotel/protobufs --go-grpc_opt=paths=source_relative ./hotel/protobufs/*.proto
		protoc -I=./bank/protobufs --go_out=./bank/protobufs --go_opt=paths=source_relative  --go-grpc_out=./bank/protobufs --go-grpc_opt=paths=source_relative ./bank/protobufs/*.proto

gen-vehicle-proto:
		protoc -I=./vehicle/api/pb --go_out=./vehicle/api/pb --go_opt=paths=source_relative  --go-grpc_out=./vehicle/api/pb --go-grpc_opt=paths=source_relative ./vehicle/api/pb/*.proto  

gen-transport-proto:
		protoc -I=./transport/api/pb --go_out=./transport/api/pb --go_opt=paths=source_relative  --go-grpc_out=./transport/api/pb --go-grpc_opt=paths=source_relative ./transport/api/pb/*.proto  
		protoc -I=./transport/pkg/adapters/storage/bank-pb --go_out=./transport/pkg/adapters/storage/bank-pb --go_opt=paths=source_relative  --go-grpc_out=./transport/pkg/adapters/storage/bank-pb --go-grpc_opt=paths=source_relative ./transport/pkg/adapters/storage/bank-pb/*.proto  
		protoc -I=./transport/pkg/adapters/storage/vehicle-pb --go_out=./transport/pkg/adapters/storage/vehicle-pb --go_opt=paths=source_relative  --go-grpc_out=./transport/pkg/adapters/storage/vehicle-pb --go-grpc_opt=paths=source_relative ./transport/pkg/adapters/storage/vehicle-pb/*.proto  
		