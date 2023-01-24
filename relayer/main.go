package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"relayer/api"
)

func main() {
	h, err := api.NewAPI()
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.HandleFunc("/permit", h.PermitHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/deposit", h.DepositWithSig).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/withdraw", h.WithdrawWithSig).Methods(http.MethodPost, http.MethodOptions)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("failed to run server: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Listen to server 8080")
}
