package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var RELAYER_PRIVATE_KEY = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

func createTransactionOpt(ctx context.Context, ec *ethclient.Client) (*bind.TransactOpts, error) {
	// Create Signer Wallet
	signerPrivateKey, err := crypto.HexToECDSA(RELAYER_PRIVATE_KEY)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	publicKey := signerPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	signerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Create Transaction options
	chainId, err := ec.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	nonce, err := ec.PendingNonceAt(ctx, signerAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(signerPrivateKey, chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000) // TODO
	auth.GasPrice = gasPrice
	return auth, nil
}

func sigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func stringToByte32(str string) [32]byte {
	hexString := strings.Replace(str, "0x", "", 1)
	decodedByteArray, _ := hex.DecodeString(hexString)
	var result [32]byte
	copy(result[:], decodedByteArray[:32])
	return result
}
