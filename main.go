package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

const msg = `
defined table users

CREATE TABLE users (
  uuid uuid NOT NULL,
  name varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  email varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
)
`

func main() {
	llm, err := openai.NewChat()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	completion, err := llm.Call(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Content: msg},
		schema.HumanChatMessage{Content: "Please generate 10 fixtures for the users table in .yaml format"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}
