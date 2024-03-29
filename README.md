# Gemini Webpage Summarier

A webpage summarizer based on google's gemini API  

## Usage
In your code
```go
import github.com/SHA-4096/gemini-webpage-summarizer/websum
```

Download this package by `go mod`
```bash
go mod tidy
```

Initialize the gemini chat session or use your existing client,session and context:  
```go
InitWebsumWithGeminiApiKey(YOUR_API_KEY)
```
```go
func ImportExistingClientAndChatSession(client,chatSession,context)
```

Use `GetWebpageSummary()` to get a summary of a webpage

```go
summary,err := GetWebPageSummary(url)
```