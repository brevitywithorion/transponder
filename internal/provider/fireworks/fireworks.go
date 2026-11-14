// Package fireworks configures the OpenAI adapter for Fireworks AI.
package fireworks

import (
	"net/http"

	"github.com/raven-clown/transponder/internal/provider/openai"
)

// DefaultBaseURL is Fireworks AI's OpenAI-compatible API, from
// https://docs.fireworks.ai/tools-sdks/openai-compatibility.
const DefaultBaseURL = "https://api.fireworks.ai/inference/v1"

// New returns an adapter for Fireworks AI. An empty baseURL means DefaultBaseURL,
// and an empty apiKey sends no Authorization header.
func New(name, baseURL, apiKey string, client *http.Client) *openai.Adapter {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return openai.New(openai.Options{Name: name, BaseURL: baseURL, APIKey: apiKey, Client: client})
}
