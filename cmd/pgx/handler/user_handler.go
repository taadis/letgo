// handler.go
package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/taadis/letgo/cmd/pgx/store"
)

type UserHandler struct {
	userStore store.UserStore
}

func NewHandler(userStore store.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (h *UserHandler) QueryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rets, err := h.userStore.QueryData(ctx)
	if err != nil {
		http.Error(w, "database query failed", http.StatusInternalServerError)
		return
	}
	slog.Info("query data success", "rets", rets)
	fmt.Fprintf(w, "query data success")
}
