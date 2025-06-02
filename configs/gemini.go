package configs

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func Gemini(ctx context.Context) *genai.Client {
	err := godotenv.Load(".env") // Load the .env file
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    geminiAPIKey := os.Getenv("GEMINI_API_KEY")
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  geminiAPIKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Fatal(err)
    }

	return client
}