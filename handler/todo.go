package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

// *HealthzHandler と同様に *TODOHandler にも ServeHTTP メソッドを実装し、HTTP のリクエストを受け取れるようにする
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := &model.TODO{}

	//HTTPメソッドがPostの場合を判定
	if r.Method == http.MethodPost {
		//CreateTODORequestにJSON Decodeを行う
		var req model.CreateTODORequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			// デコードエラーの場合、400 Bad Request を返す
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		//subjectが空文字列の場合をif文で判定し、空の場合は400 BadRequestとしてHTTPResponseを返す
		if req.Subject == "" {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		//r.Context()を呼びだして、リクエストのコンテキストを取得
		ctx := r.Context()
		// CreateTODOメソッドを呼び出してDBに保存
		todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)

		if err != nil {
			// エラーが発生した場合、500 Internal Server Error を返す
			http.Error(w, "Failed to create TODO", http.StatusInternalServerError)
			return
		}

		// 保存したTODOをCreateTODOResponseに代入
		resp := &model.CreateTODOResponse{TODO: *todo}

		// JSON Encodeを行いHTTP Responseを返す
		//HTTPレスポンスのヘッダーにContent-Typeを設定するのはクライアントがレスポンスを正しく解釈できるようにするため
		//JSON形式のデータを送信する場合、Content-Typeを"application/json"に設定するのが一般的
		w.Header().Set("Content-Type", "application/json")
		// 成功時のステータスコードを200 OKに設定
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

		return
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		return
	}
}
