package main

import (
	"bytes"         //リクエストのデータをバイト形式で保持す
	"encoding/json" //Goの構造体をJSONにエンコードしたり、JSONをGoの構造体にデコードするため
	"fmt"           //フォーマットした文字列を出力する用
	"log"           //エラーメッセージの出力用
	"net/http"      //APIリクエストを送信する用
	"os"            //環境変数を扱うためのパッケージ

	"github.com/joho/godotenv"
)

// OpenAI APIのエンドポイントを設定
const (
	openAIAPIURL = "https://api.openai.com/v1/chat/completions" // APIエンドポイント
)

// APIリクエストのJSON構造を表す リクエストにはモデル名（Model）とメッセージ（Messages）の配列が含まれる
type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// 一つ一つのメッセージを表す　役割と内容が含まれる
type ChatMessage struct {
	Role    string `json:"role"` // "user" または "assistant"
	Content string `json:"content"`
}

// APIからのレスポンス全体(複数の応答)を表す
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

// OpenAIの返答として生成される各応答メッセージを保持
type Choice struct {
	Message ChatMessage `json:"message"`
}

// OpenAI APIにリクエストを送信し、レスポンスを取得する
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

	// リクエストデータの構築 (OpenAIRequest型で)
	reqBody := OpenAIRequest{
		Model:    model,
		Messages: messages,
	}

	// reqBodyをJSON形式にエンコード　→ APIに送信できる状態にする
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to encode request: %v", err)
	}

	// POSTメソッドでAPIエンドポイントにHTTPリクエストの作成
	req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// HTTPヘッダーの設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// HTTPクライアントの作成
	client := &http.Client{}

	// HTTPリクエストの送信
	resp, err := client.Do(req)

	// リクエストを実行し、APIの応答を受け取る　応答の後にdeferでクローズ
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// ステータスコードのチェック
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// レスポンスをデコード → OpenAIResponseに格納
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
