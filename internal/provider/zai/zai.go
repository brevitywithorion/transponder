// Package zai configures the OpenAI adapter for Z.ai (Zhipu GLM).
package zai

import (
	"net/http"

	"github.com/raven-clown/transponder/internal/provider/openai"
)

// DefaultBaseURL is Z.ai's OpenAI-compatible API, from
// https://docs.z.ai/guides/overview/quick-start. The path ends in /v4, and /chat/completions is appended as given. The China host is https://open.bigmodel.cn/api/paas/v4.
const DefaultBaseURL = "https://api.z.ai/api/paas/v4"

// New returns an adapter for Z.ai (Zhipu GLM). An empty baseURL means DefaultBaseURL,
// and an empty apiKey sends no Authorization header.
func New(name, baseURL, apiKey string, client *http.Client) *openai.Adapter {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return openai.New(openai.Options{Name: name, BaseURL: baseURL, APIKey: apiKey, Client: client})
}
