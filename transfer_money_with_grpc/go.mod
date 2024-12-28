module main

go 1.22.5

replace github.com/onlineTraveling/bank => ../bank

require (
	github.com/onlineTraveling/bank v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.69.2
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.36.0 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
