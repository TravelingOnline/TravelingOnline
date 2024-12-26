module main

go 1.22.5

replace github.com/onlineTraveling/bank => ../bank

require (
	github.com/google/uuid v1.6.0
	github.com/streadway/amqp v1.1.0
)
