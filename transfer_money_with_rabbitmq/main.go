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
	SenderOwnerID   string `json:"wallet_id_from"`
	ReceiverOwnerID string `json:"wallet_id_to"`
	Amount          int    `json:"amount"`
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
		SenderOwnerID:   "e5824ec0-48ae-4712-9178-48f96e33328c",
		ReceiverOwnerID: "f2fc63c0-443e-4181-b0c9-2316a4a9845c",
		Amount:          120,
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
