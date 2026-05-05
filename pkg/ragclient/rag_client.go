package ragclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type RAGClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewRAGClient() *RAGClient {
	url := os.Getenv("RAG_SERVICE_URL")
	if url == "" {
		url = "http://rag-service:8001"
	}
	return &RAGClient{
		BaseURL:    url,
		HTTPClient: &http.Client{Timeout: 90 * time.Second},
	}
}

func (c *RAGClient) post(endpoint string, payload interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Post(c.BaseURL+endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("rag request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errRes map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errRes)
		return nil, fmt.Errorf("rag service error (status %d): %v", resp.StatusCode, errRes)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("rag decode error: %w", err)
	}
	return result, nil
}

func (c *RAGClient) Generate(payload map[string]interface{}) (map[string]interface{}, error) {
	return c.post("/rag/generate", payload)
}

func (c *RAGClient) Refine(payload map[string]interface{}) (map[string]interface{}, error) {
	return c.post("/rag/refine", payload)
}

func (c *RAGClient) Ask(payload map[string]interface{}) (map[string]interface{}, error) {
	return c.post("/rag/ask", payload)
}
