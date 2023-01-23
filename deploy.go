package main

import (
	"fmt"
	"strings"
	"context"
	"time"
	"log"
	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	store "./contracts" // for demo
	functions "./func"
)

var pairing *pbc.Pairing
var g *pbc.Element

func CommContract(paramsByte []byte) {

	gethPath, nodeKey := functions.GethPathAndKey()
    
	connection, err := ethclient.Dial(gethPath)
	functions.CheckError(err)
	fmt.Println("connection Returned: ", connection)
	ctx := context.Background()

	auth, err := bind.NewTransactor(strings.NewReader(nodeKey), "account1")
	functions.CheckError(err)

	start := time.Now()
	addr, tx, instContract, err := store.DeployStore(auth, connection)
	functions.CheckError(err)
	fmt.Println("\n addr Returned: ", addr)
	fmt.Println("\n instContract Returned: ", instContract)

	// Write contract address to file for other app reference
	functions.FileCrWr("./tmp/contractInfo/contract-address", []byte(addr.Hex()))

	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)
	fmt.Println("tx=",tx.Hash().Hex())
	elapsed := time.Since(start)
    	log.Printf("\nContract deployment took %s \n", elapsed)

	y := func() {
        defer timeTrack(time.Now(), "y")
	pairing, err = pbc.NewPairingFromString(string(paramsByte))
	functions.CheckError(err)
	g = pairing.NewG1().Rand()
	}
	y()
	
	z := func() {
        defer timeTrack(time.Now(), "z")

	tx1, err := instContract.SetParams(auth, paramsByte, g.Bytes())
	functions.CheckError(err)
	fmt.Println("\ntx1=",tx1.Hash().Hex())
	
	_, err = bind.WaitMined(ctx, connection, tx1)
	functions.CheckError(err)
	}
	z()

}

func timeTrack(start time.Time, name string) {
        elapsed := time.Since(start)
        fmt.Printf("\nFunction %s took %s", name,elapsed)
}

func main() {

	// create ./tmp/ folder for storing secrets information of parties and public parameters 
	functions.Createtmp()
	params := pbc.GenerateA(160, 512)
	CommContract([]byte(params.String()))
}

