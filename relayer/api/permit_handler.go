package api

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"log"
	"math/big"
	"net/http"
)

type PermitRequestParams struct {
	Owner     common.Address `json:"owner"`
	Spender   common.Address `json:"spender"`
	Value     int64          `json:"value"`
	DeadLine  int64          `json:"deadline"`
	Signature string         `json:"signature"`

	V uint8  `json:"v"`
	R string `json:"r"`
	S string `json:"s"`
}

func (h *Handler) PermitHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var params PermitRequestParams
	if err := json.Unmarshal(body, &params); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ctx := context.Background()
	auth, err := createTransactionOpt(ctx, h.EthClient)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("Signature: %v\n", params.Signature)
	log.Printf("params.Owner: %v\n", params.Owner)
	log.Printf("params.Spender: %v\n", params.Spender)
	log.Printf("value: %v\n", params.Value)
	log.Printf("deadline: %v\n", params.DeadLine)
	permitTx, err := h.Erc20Token.Permit(
		auth,
		params.Owner,
		common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8"),
		big.NewInt(params.Value),
		big.NewInt(params.DeadLine),
		params.V,
		stringToByte32(params.R),
		stringToByte32(params.S),
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}

	transferAuth, err := createTransactionOpt(ctx, h.EthClient)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}
	transferTx, err := h.Erc20Token.Transfer(transferAuth, common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"), big.NewInt(params.Value))
	if err != nil {
		log.Println("Transfer:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("transferTx: %+v\n", transferTx)

	log.Printf("h.Erc20Token: %+v\n", h.Erc20Token)
	balance, err := h.Erc20Token.BalanceOf(&bind.CallOpts{}, params.Owner)
	if err != nil {
		log.Println("BalanceOf:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("balance: %+v\n", balance)

	resp, err := json.Marshal(permitTx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
