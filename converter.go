package main

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"os"
	"strings"
)

const model = "gemini-2.0-flash"

func convert(ctx context.Context, apiKey, filename, outname string) error {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return fmt.Errorf("failed to create gemini client:%w", err)
	}

	pdfBytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read input file:%w", err)
	}

	req := []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					InlineData: &genai.Blob{
						MIMEType: "application/pdf",
						Data:     pdfBytes,
					},
				},
				{
					Text: "convert this pdf to markdown",
				},
			},
		},
	}

	resp, err := client.Models.GenerateContent(ctx, model, req, nil)
	if err != nil {
		return fmt.Errorf("failed to generate content with gemini:%w", err)
	}

	file, err := os.Create(outname)
	if err != nil {
		return fmt.Errorf("failed to create output file:%w", err)
	}
	defer file.Close()

	_, err = file.WriteString(cleanMarkdownBlock(resp.Text()))
	if err != nil {
		return fmt.Errorf("failed to write to output file:%w", err)
	}

	return nil
}

func cleanMarkdownBlock(markdown string) string {
	md := strings.TrimPrefix(markdown, "```markdown\n")
	md = strings.TrimSuffix(md, "```")
	return strings.TrimSpace(md)
}
