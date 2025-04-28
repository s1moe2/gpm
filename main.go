package main

import (
	"context"
	"flag"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Usage: gpm <pdf-file> <output-md-file>")
	}

	if err := convert(ctx, apiKey, args[0], args[1]); err != nil {
		log.Fatalf("Error converting PDF to Markdown: %v", err)
	}
}
