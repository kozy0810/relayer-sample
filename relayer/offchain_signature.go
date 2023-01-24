package main

//import (
//	"crypto/ecdsa"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/common/math"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/signer/core/apitypes"
//	"log"
//)
//
//type MessageParams struct {
//	Owner    string
//	Spender  string
//	Value    *math.HexOrDecimal256
//	Nonce    *math.HexOrDecimal256
//	Deadline *math.HexOrDecimal256
//}
//
//var typesStandard = apitypes.Types{
//	"EIP712Domain": {
//		{
//			Name: "name",
//			Type: "string",
//		},
//		{
//			Name: "version",
//			Type: "string",
//		},
//		{
//			Name: "chainId",
//			Type: "uint256",
//		},
//		{
//			Name: "verifyingContract",
//			Type: "address",
//		},
//	},
//	"Permit": {
//		{
//			Name: "owner",
//			Type: "address",
//		},
//		{
//			Name: "spender",
//			Type: "address",
//		},
//		{
//			Name: "value",
//			Type: "uint256",
//		},
//		{
//			Name: "nonce",
//			Type: "uint256",
//		},
//		{
//			Name: "deadline",
//			Type: "uint256",
//		},
//	},
//}
//
//const primaryType = "Permit"
//
//func getDomain() apitypes.TypedDataDomain {
//	return apitypes.TypedDataDomain{
//		Name:              "MyToken",
//		Version:           "1",
//		ChainId:           math.NewHexOrDecimal256(31337),
//		VerifyingContract: ERC20_ADDRESS,
//		//Salt:              "salt",
//	}
//}
//
//func generateMessage(params MessageParams) apitypes.TypedDataMessage {
//	return apitypes.TypedDataMessage{
//		"owner":    params.Owner,
//		"spender":  params.Spender,
//		"value":    params.Value,
//		"nonce":    params.Nonce,
//		"deadline": params.Deadline,
//	}
//}
//
//type headlessUi struct {
//	approveCh chan string // to send approve/deny
//	inputCh   chan string // to send password
//}
//
//func GenerateSig(params MessageParams, privateKey *ecdsa.PrivateKey) (hexutil.Bytes, error) {
//	domainStandard := getDomain()
//	messageStandard := generateMessage(params)
//
//	message := apitypes.TypedData{
//		Types:       typesStandard,
//		PrimaryType: primaryType,
//		Domain:      domainStandard,
//		Message:     messageStandard,
//	}
//	log.Printf("message: %+v\n", message)
//
//	//db, _ := fourbyte.New()
//	//api := core.NewSignerAPI(nil, 31337, true, nil, db, true, &storage.NoStorage{})
//	//log.Printf("api: %+v\n", api)
//	//
//	//resp, err := api.SignTypedData(context.Background(), *signer, message)
//
//	return crypto.Sign(message.TypeHash("Permit"), privateKey)
//	//if err != nil {
//	//	log.Printf("Error: %+v\n", err)
//	//}
//	//log.Printf("resp: %+v\n", resp)
//	//return nil, nil
//}
