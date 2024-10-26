package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// OpenAI APIのエンドポイントとAPIキーを設定
const (
	openAIAPIURL = "https://api.openai.com/v1/chat/completions" // APIエンドポイント
)

// OpenAIのリクエスト構造体
type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// ChatMessage構造体は、メッセージを表す
type ChatMessage struct {
	Role    string `json:"role"` // "user" または "assistant"
	Content string `json:"content"`
}

// OpenAIResponse構造体は、APIからのレスポンスを表す
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

// Choice構造体は、選択肢を表す
type Choice struct {
	Message ChatMessage `json:"message"`
}

// OpenAIAPI関数は、OpenAI APIにリクエストを送信し、レスポンスを取得する
func OpenAIAPI(model string, messages []ChatMessage) (string, error) {

	// .envファイルを読み込む
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// API_KEYを環境変数から取得
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set in .env file")
	}

	// リクエストの構築
	reqBody := OpenAIRequest{
		Model:    model,
		Messages: messages,
	}

	// JSON形式にエンコード
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to encode request: %v", err)
	}

	// HTTPリクエストの作成
	req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// ヘッダーの設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// ステータスコードのチェック
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// レスポンスをデコード
	var openAIResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// レスポンスからメッセージを取得
	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no choices returned from OpenAI API")
}
