package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"w4ll3t/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type mockService struct{}

func (m *mockService) Update(ctx context.Context, id uuid.UUID, amount int64, op string) (int64, error) {
	return 1000, nil
}

func (m *mockService) Get(ctx context.Context, id uuid.UUID) (int64, error) {
	return 1000, nil
}

func TestHTTP_UpdateWallet(t *testing.T) {
	mock := &mockService{}
	h := handler.NewWalletHandler(mock)

	r := chi.NewRouter()
	r.Post("/api/v1/wallet", h.UpdateWallet)

	body := map[string]interface{}{
		"walletId":      uuid.New(),
		"operationType": "DEPOSIT",
		"amount":        1000,
	}

	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
