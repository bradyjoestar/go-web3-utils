package main

import (
	"fmt"
	"github.com/onrik/ethrpc"
	"os"
	"strconv"
)

func main() {
	fmt.Println("wenbin test")

	client := ethrpc.New("http://127.0.0.1:8545")

	//
	//fmt.Println("--------------------getTransction        --------------------------")
	//getTransaction(os.Args[1])
	//
	//fmt.Println("\n\n\n\n\n")
	//fmt.Println("--------------------getTransactionReceipt--------------------------")
	//getTransactionReceipt(os.Args[1])

	fmt.Println("----------------getLastesBlocks-------------------------------------")
	blockNumber, err := client.EthBlockNumber()
	fmt.Println(blockNumber)

	for i := blockNumber; i > blockNumber-3; i-- {
		blocks, _ := client.EthGetBlockByNumber(i, true)
		fmt.Println("===============block Number================================")
		fmt.Println(blocks.Hash)
		fmt.Println(blocks.Number)
		fmt.Println(blocks.Timestamp)
		fmt.Println("---------------transaction details-------------------------")
		fmt.Println("hash: " + blocks.Transactions[0].Hash)
		fmt.Println("from: " + blocks.Transactions[0].From)
		fmt.Println("to: " + blocks.Transactions[0].To)
		fmt.Println("data: " + blocks.Transactions[0].Input)
		fmt.Println(blocks.Transactions[0].Nonce)
	}

}

func getTransaction(args string) {
	client := ethrpc.New("http://127.0.0.1:8545")

	tx, err := client.EthGetTransactionByHash(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("from: " + tx.From)
	fmt.Println("to: " + tx.To)
	fmt.Println("nonce: " + strconv.Itoa(tx.Nonce))
	fmt.Println("txHash: " + tx.Hash)
	fmt.Println("input data: " + tx.Input)
}

func getTransactionReceipt(args string) {
	client := ethrpc.New("http://127.0.0.1:8545")

	receipt, err := client.EthGetTransactionReceipt(args)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(receipt)
	fmt.Println("hash: " + receipt.TransactionHash)
	fmt.Println("status: " + receipt.Status)
	fmt.Println("contract address: " + receipt.ContractAddress)
	for _, v := range receipt.Logs {
		fmt.Println("logs: ")
		fmt.Println(v)
	}
}
