package main

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration struct to hold bot token
type Configuration struct {
	TelegramBotToken string `json:"telegram_bot_token"`
}

func getToken() string {

	// Open and read the configuration file
	configFile, err := os.Open("D:\\Projects\\CopilotTelegramBot\\config.json")
	if err != nil {
		log.Fatal("Error opening configuration file:", err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Fatal("Error closing configuration file:", err)
		}
	}(configFile)

	// Decode the JSON configuration
	var config Configuration
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal("Error decoding configuration:", err)
	}

	// Access the bot token from the configuration
	return config.TelegramBotToken
}
