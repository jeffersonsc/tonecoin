package main

import (
	"flag"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/novatrixtech/tonecoin/proto"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "mining new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("192.168.0.104:8097", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to the server. Error: %s\n", err.Error())
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		listBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("Unable to add block. Error: %s\n", err.Error())
	}
	log.Printf("%s - New block generated: %s", time.Now().String(), block.Hash)
}

func listBlockchain() {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("Unable to get the blockchain. Error: %s\n", err.Error())
	}
	log.Println("Blocks: ")
	for _, b := range bc.Blocks {
		log.Printf("Hash: %s - Previous hash: %s - Data: %+v\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
