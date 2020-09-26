package handlers

import (
	"context"
	"log"
	"math/big"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/phoax/godemo/models"
	op "github.com/phoax/godemo/restapi/operations/account"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func getClient() (client *ethclient.Client, err error) {
// 	log.Println("Connected to node")
// 	client, err = ethclient.Dial(apiconfig.GetString("ETHEREUM_NODE_URL"))
// 	return
// }
func GetNonce(client *ethclient.Client, address string) uint64 {
	from := common.HexToAddress(address)

	// Get context
	ctx := context.Background()

	// Get nonce
	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		log.Printf("Nonce fail", err)
	}
	return nonce
}

// Transfer handler
func AccountTransferHandler() middleware.Responder {
	// Connect to node
	client, err := getClient()
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}

	fromAddress := "0xfdFa0f372a909F33945a4Eb207Fd0Df694D02570"
	privateKey, _ := crypto.HexToECDSA("df155959a1d1a37caae2087eee40f934eafc0a1f4a7865785f169e5871dbf765")
	to := common.HexToAddress("0x940a7acD2A11846ba2F5Fc52f68EC69daFe8C550")
	nonce := GetNonce(client, fromAddress)
	log.Println("Current Nonce: ", nonce)

	ctx := context.Background()

	// Figure out the gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Gas price: ", gasPrice)

	// Figure out the gas limit
	from := common.HexToAddress(fromAddress)

	value := big.NewInt(2000000000000000000) // 2 ether

	input := []byte("")
	msg := ethereum.CallMsg{From: from, To: &to, Value: value, Data: input}
	gasLimit, err := client.EstimateGas(ctx, msg)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Gas limit: ", gasLimit)

	log.Printf("Create TX")

	// Create the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, input)

	log.Printf("TX OK: %v", privateKey)

	// Sign
	transact_opt := bind.NewKeyedTransactor(privateKey)

	log.Printf("transact_opt: %T", transact_opt)

	tx, err = transact_opt.Signer(types.HomesteadSigner{}, common.HexToAddress(fromAddress), tx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Tx: %v", tx)

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Printf("Transaction failed: %v", err)
	}

	//
	payload := models.Ack{Status: "success", Message: "transfer ok"}
	return op.NewSetTransferCreated().WithPayload(&payload)
}

func AccountBalanceHandler() middleware.Responder {
	client, err := getClient()

	// ethereum.GetBalance(client)
	to := common.HexToAddress("0x940a7acD2A11846ba2F5Fc52f68EC69daFe8C550")

	ctx := context.Background()
	balance, err := client.BalanceAt(ctx, to, nil)
	if err != nil {
		log.Printf("could not get balance: %v", err)
	}

	// Check balance
	log.Println("Balance", balance)

	payload := models.Ack{Status: "success", Message: "balance ok"}
	return op.NewSetTransferCreated().WithPayload(&payload)
}
