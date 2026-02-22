package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"w4ll3t/internal/model"
	"w4ll3t/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type WalletService interface {
	Update(ctx context.Context, id uuid.UUID, amount int64, op string) (int64, error)
	Get(ctx context.Context, id uuid.UUID) (int64, error)
}

type WalletHandler struct {
	service WalletService
}

func NewWalletHandler(s WalletService) *WalletHandler {
	return &WalletHandler{service: s}
}
func (h *WalletHandler) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	var req model.WalletOperation
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	balance, err := h.service.Update(
		r.Context(),
		req.WalletID,
		req.Amount,
		string(req.OperationType),
	)

	if err != nil {
		switch err {
		case repository.ErrInsufficientFunds:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case repository.ErrWalletNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"balance": balance})
}

func (h *WalletHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	balance, err := h.service.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"balance": balance})
}
