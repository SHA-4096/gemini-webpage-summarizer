package websum

import (
	"io"
	"net/http"
)

func InitWebsumWithGeminiApiKey(apiKey string) {
	InitGemini(apiKey)
}

func GetWebpageSummary(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "create request failed", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "make request failed", err
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "reading the response body failed", err
	}
	// Convert the body to a string
	bodyStr := string(body)
	// Chat with Gemini
	ret, err := ChatWithGemini(bodyStr)
	return ret, err
}
