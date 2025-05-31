package library

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func Gemini(content string, highlighter string) string {
	err := godotenv.Load(".env") // Load the .env file
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    geminiAPIKey := os.Getenv("GEMINI_API_KEY")
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  geminiAPIKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Fatal(err)
    }

	prompt := `
        You are a helpful assistant.

        The user highlighted the following text from a website:
        "%s"

        Here is the surrounding context from the page:
        "%s"

        Instructions:
        - If the highlighted text is a question, respond to it as if the question were directed at you. Answer in the voice of the person you're representing.
        - If it's not a question, provide a brief explanation and ask if the user needs any further assistance.
    `

    highlight := highlighter  // "I think this is safe to integrate. Could you update this PR branch?" // question
    fmt.Println("highlighted", highlighter)
    finalPrompt := fmt.Sprintf(prompt, highlight, content)

    color.Yellow("PROMPT\n")
    fmt.Println(finalPrompt)

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.0-flash",
        genai.Text(finalPrompt),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    color.Blue("RESULT\n")
    fmt.Println(result.Text())
	return result.Text()
}