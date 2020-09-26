package handlers

import (
	"context"
	"log"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/phoax/godemo/models"
	op "github.com/phoax/godemo/restapi/operations/network"
)

func getClient() (client *ethclient.Client, err error) {
	log.Println("Connected to node")
	client, err = ethclient.Dial(apiconfig.GetString("ETHEREUM_NODE_URL"))
	return
}

// BlockInfo handler
func BlockNumberHandler() middleware.Responder {
	// Connect to node
	client, err := getClient()
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}
	// Get context
	ctx := context.Background()
	// Block number
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(header.Number.ToString())
	// Response
	payload := models.NetworkBlockNumber{BlockNumber: header.Number.Int64()}
	return op.NewGetBlockNumberOK().WithPayload(&payload)
}
