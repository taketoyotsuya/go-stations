package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	//構造体の中の変数のことをFeildという
	//`Message`フィールドはJSONのキー名が`message`になる
	Message string `json:"message"`
}
