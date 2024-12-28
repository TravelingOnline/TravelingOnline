package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// CreateWalletRequest defines the structure of the request message
type CreateWalletRequest struct {
	UserID string `json:"user_id"`
}
type TransferRequest struct {
	WalletIDFrom string `json:"wallet_id_from"`
	WalletIDTo   string `json:"wallet_id_to"`
	Amount       int    `json:"amount"`
}

// Function to publish a message to RabbitMQ
func publishToRabbitMQ(queueName string, message []byte) error {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Update with your RabbitMQ credentials
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue (idempotent)
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-deleted
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return err
	}

	// Publish the message
	err = ch.Publish(
		"",        // Exchange
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message published to RabbitMQ: %s", string(message))
	return nil
}

func main() {

	request := TransferRequest{
		WalletIDFrom: "bd8873c8-be31-4285-b8ef-b7ed2ac82be8",
		WalletIDTo:   "5f664e09-e970-48c2-8fc2-dceb3096bd52",
		Amount:       120,
	}
	// Serialize the message to JSON
	message, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Failed to serialize request: %v", err)
	}
	queueName := "bank-service/transfer-transaction"
	err = publishToRabbitMQ(queueName, message)
	if err != nil {
		log.Fatalf("Error publishing to RabbitMQ: %v", err)
	}

	log.Println("Transfer message sent to RabbitMQ successfully")

}
