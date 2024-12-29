package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// import (
// 	"encoding/json"
// 	"os"
// )

// func ReadConfig(configPath string) (Config, error) {
// 	var c Config
// 	data, err := os.ReadFile(configPath)
// 	if err != nil {
// 		return c, err
// 	}

// 	return c, json.Unmarshal(data, &c)

// }

// func MustReadConfig(configPath string) Config {
// 	c, err := ReadConfig(configPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c
// }

func LoadConfig() (Config, error) {
	var cfg Config

	// Load .env file into environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables directly")
	}

	// Bind environment variables with Viper
	viper.AutomaticEnv()

	// Read and map the database configuration
	cfg.DB.Host = viper.GetString("DB_HOST")
	cfg.DB.Port = getIntFromEnv("DB_PORT", 5434) // Default to 5432 if not set
	cfg.DB.Database = viper.GetString("DB_DATABASE")
	cfg.DB.User = viper.GetString("DB_USER")
	cfg.DB.Password = viper.GetString("DB_PASSWORD")

	return cfg, nil
}

// Helper function to convert environment variable to integer
func getIntFromEnv(key string, defaultValue int) int {
	val := viper.GetString(key)
	if val == "" {
		return defaultValue
	}
	parsedVal, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Invalid integer value for %s: %v, using default: %d\n", key, err, defaultValue)
		return defaultValue
	}
	return parsedVal
}
