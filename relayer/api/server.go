package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"net/http"
	erc20Token "relayer/contracts/ERC20Token"

	"github.com/gorilla/mux"
)

const (
	DefaultServerPort = "8080"
)

type Handler struct {
	EthClient  *ethclient.Client
	Erc20Token *erc20Token.Erc20Token
}

func NewHandler() *Handler {
	ec, err := ethclient.Dial("http://127.0.0.1:8545")
	erc20TokenInstance, err := erc20Token.NewErc20Token(common.HexToAddress("0x3F71a85B94314b54889A8AC050C571dC5406797A"), ec)
	if err != nil {
		log.Fatalln(err)
	}

	return &Handler{
		EthClient:  ec,
		Erc20Token: erc20TokenInstance,
	}
}

func NewAPI() (*Handler, error) {
	return NewHandler(), nil
}

func (h *Handler) NewServe(port string, handler http.Handler) error {
	if port == "" {
		port = DefaultServerPort
	}

	return http.ListenAndServe(":"+port, handler)
}

func (h *Handler) registerHandler() *mux.Router {
	r := mux.NewRouter()
	r.Name("Hoge").Subrouter()
	return r
}

type Services struct{}

func NewServices() (*Services, error) {
	//var err error
	s := new(Services)

	return s, nil
}
