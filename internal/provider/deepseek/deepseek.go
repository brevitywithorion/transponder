// Package deepseek configures the OpenAI adapter for DeepSeek.
package deepseek

import (
	"net/http"

	"github.com/raven-clown/transponder/internal/provider/openai"
)

// DefaultBaseURL is DeepSeek's OpenAI-compatible API, from
// https://api-docs.deepseek.com/. Chat is POST /chat/completions on this host.
const DefaultBaseURL = "https://api.deepseek.com"

// New returns an adapter for DeepSeek. An empty baseURL means DefaultBaseURL,
// and an empty apiKey sends no Authorization header.
func New(name, baseURL, apiKey string, client *http.Client) *openai.Adapter {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return openai.New(openai.Options{Name: name, BaseURL: baseURL, APIKey: apiKey, Client: client})
}
