package helper

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/pkg/adapters/clients/grpc/mappers"
	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc"
)

func EncryptAES(plaintext string, key []byte) (string, error) {
	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt and concatenate nonce
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Return base64 encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES decrypts string with AES-GCM
func DecryptAES(ciphertext string, key []byte) (string, error) {
	// Decode base64
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract nonce size
	nonceSize := gcm.NonceSize()
	if len(decoded) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertextBytes := decoded[:nonceSize], decoded[nonceSize:]

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// GenerateRandomKey generates a random key of specified size
func GenerateRandomKey(size int) ([]byte, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
func CreateWalletGrpc(cntx context.Context, userID uuid.UUID) error {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client cannot connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new BankService client
	client := protobufs.NewBankServiceClient(conn)

	// Prepare the request
	in := &protobufs.CreateWalletRequest{
		UserID: userID.String(),
	}

	// Call the CreateWallet method
	response, err := client.CreateWallet(cntx, in)
	if err != nil {
		log.Fatalf("cannot create wallet: %v", err)
	}

	// Map the response to the domain model
	domainResponse, err := mappers.CreateWalletResponseToMessageDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)

		return err
	}

	// Print the domain response
	log.Printf("Wallet created successfully: %v", *domainResponse)
	return nil
}
