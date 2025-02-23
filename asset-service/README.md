# Asset Service
This service maintains and monitors key asset information, including:
- Valuation - Current market worth and financial metrics
- Future Growth - Projected performance and industry trends
- Past Performance - Historical financial and operational results
- Financial Health - Stability, liquidity, and debt analysis
- Dividend - Payout history and sustainability
- Management - Leadership structure and governance
- Ownership - Shareholder composition and insider activity
- News and public sentiment - Media coverage and market perception


## AI Generated Content
Here's an idea on how I will use Open AI and Redis to generate content for a company.

```go
package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sashabaranov/go-openai"
)

// Initialize Redis client
var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Change if needed
})

// OpenAI API Key (use environment variable in production)
var openaiClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

// Hash function for caching
func hashPrompt(prompt string) string {
	hash := sha256.Sum256([]byte(prompt))
	return hex.EncodeToString(hash[:])
}

// Function to get response from cache or OpenAI API
func getOpenAIResponse(prompt string) (string, error) {
	// Generate hash key for Redis
	cacheKey := hashPrompt(prompt)

	// Check Redis cache
	cachedResponse, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		return cachedResponse, nil // Cache hit
	}

	// If cache miss, fetch from OpenAI API
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a financial analyst providing detailed and accurate stock analysis."},
			{Role: "user", Content: prompt},
		},
	}
	resp, err := openaiClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	responseText := resp.Choices[0].Message.Content

	// Store response in Redis with a 24-hour expiration
	err = redisClient.Set(ctx, cacheKey, responseText, 24*time.Hour).Err()
	if err != nil {
		log.Println("Failed to cache response:", err)
	}

	return responseText, nil
}

// Main function to start server
func main() {
	prompt := "Summarize the latest news and sentiment analysis for NVIDIA (NVDA) stock. Include key headlines and their" +
		" potential impact on the stock price."

	response, err := getOpenAIResponse(prompt)
	if err != nil {
		log.Println("Error getting response:", err)
		return
	}

	log.Println(response)

}

```


