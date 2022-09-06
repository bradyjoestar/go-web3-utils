package main

import (
	"fmt"
	"github.com/onrik/ethrpc"
	"log"
	"os"
)

func main() {
	fmt.Println("wenbin test")

	client := ethrpc.New("http://127.0.0.1:8545")

	fmt.Println(os.Args[1])

	version, err := client.Web3ClientVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)

	if len(os.Args) != 2 {
		panic("txhash should be provided")
	}

	tx, err := client.EthGetTransactionByHash(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tx)
	fmt.Println("from" + tx.From)
	fmt.Println("to" + tx.To)
	fmt.Println("input data" + tx.Input)
}
