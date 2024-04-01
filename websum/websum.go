package websum

import (
	"io"
	"net/http"
	"regexp"
)

func InitWebsumWithGeminiApiKey(apiKey string, customPrompt *string) {
	if customPrompt != nil {
		InitGemini(apiKey, *customPrompt)
	} else {
		InitGemini(apiKey, "Now you're an assitant to help me with web page summary. You'll receive contents in webpage and you will help me summarize it in its original language.")
	}

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
	bodyStr = removeHtmlTags(bodyStr)
	ret, err := ChatWithGemini("Please summarize: " + bodyStr)
	return ret, err
}

func removeHtmlTags(str string) string {
	re := regexp.MustCompile(">(.*?)<")
	matches := re.FindAllString(str, -1)
	ret := ""
	for _, match := range matches {
		ret += match[1:len(match)-1] + " "
	}
	return ret
}
