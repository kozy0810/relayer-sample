package api

import "net/http"

func (h *Handler) WithdrawWithSig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WithdrawWithSig"))
}
