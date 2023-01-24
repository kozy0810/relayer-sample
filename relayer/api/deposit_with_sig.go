package api

import "net/http"

func (h *Handler) DepositWithSig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DepositWithSig"))
}
