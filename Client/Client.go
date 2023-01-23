package main

import (
	"fmt"
	"strings"
	"context"
	"io/ioutil"
	"net"
	"bufio"
	"os"
	"time"
	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	store "../contracts" // for demo
	functions "../func"
	//fp "../fairswap"
	
)

var pairing *pbc.Pairing
var pk_C *pbc.Element
var sk_C *pbc.Element
var result bool
var test_key []byte
var rct []byte

func ContractVars() (*store.Store, *bind.TransactOpts, *ethclient.Client) {
	gethPath, nodeKey := functions.GethPathAndKey()
    
	addrByte, err := ioutil.ReadFile("../tmp/contractInfo/contract-address")
	functions.CheckError(err)
    
	connection, err := ethclient.Dial(gethPath)
	functions.CheckError(err)
    
	instContract, err := store.NewStore(common.HexToAddress(string(addrByte)), connection)
	functions.CheckError(err)
    
	auth, err := bind.NewTransactor(strings.NewReader(nodeKey), "account1")
	functions.CheckError(err)

	return instContract, auth, connection
}

func tcpClient (types string, address string) string {

	conn, _ := net.Dial(types, address)
	for { 
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter which data to access: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)

		return message
	}
}

func timeTrack(start time.Time, name string) {
        elapsed := time.Since(start)
        fmt.Printf("\nFunction %s took: %s \n", name,elapsed)
}

func main() {

	instContract, auth, connection := ContractVars()
	fmt.Println("connection Returned: ", connection)
	fmt.Println("Auth Returned: ", auth)
	ctx := context.Background()

	GetParamsByte_val, GetG_val, err := instContract.GetParams(nil)
	functions.CheckError(err)
	
	pairing, err = pbc.NewPairingFromString(string(GetParamsByte_val))
	functions.CheckError(err)

	g := pairing.NewG1().SetBytes(GetG_val)
	
	v := func() {
        defer timeTrack(time.Now(), "v")
	sk_C = pairing.NewZr().Rand()
	pk_C = pairing.NewG1().PowZn(g, sk_C)
	}
	v()

	functions.FileCrWr("../tmp/private/client/skC", sk_C.Bytes())

	w := func() {
        defer timeTrack(time.Now(), "w")
	tx, err := instContract.SetPKC(auth, pk_C.Bytes())
	functions.CheckError(err)
	fmt.Println("tx=",tx.Hash().Hex())
	
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)
	}
	w()

	message:= tcpClient("tcp", "127.0.0.1:8081")

	if (strings.TrimRight(message, "\n") == "Acknowledged") {

		nonce, err := instContract.GetNonce(nil)
		functions.CheckError(err)

		time.Sleep(80 * time.Second)

		c11_val, err := ioutil.ReadFile("../tmp/private/client/SharedData/c11")
		functions.CheckError(err)
		c11 := pairing.NewG1().SetBytes(c11_val)
		c12_val, err := ioutil.ReadFile("../tmp/private/client/SharedData/c12")
		functions.CheckError(err)
		c12 := pairing.NewG1().SetBytes(c12_val)
		c2_val, err := ioutil.ReadFile("../tmp/private/client/SharedData/c2")
		functions.CheckError(err)
		c2 := pairing.NewGT().SetBytes(c2_val)

		rct = append(rct, c11.Bytes()...)
		rct = append(rct, c12.Bytes()...)
		rct = append(rct, c2.Bytes()...)
		
		x := func() {
        	defer timeTrack(time.Now(), "x")
		hash1 := crypto.Keccak256Hash(rct)
		result, err = instContract.CompareMeta2(nil, hash1.Hex())
		fmt.Println("Result=",result)
		functions.CheckError(err)
		}
		x()

		if (result == true) {
			
			y := func() {
        		defer timeTrack(time.Now(), "y")
			res := pairing.NewGT().Div(c2,pairing.NewGT().PowZn(pairing.NewGT().Pair(c11, c12), pairing.NewZr().Invert(sk_C)))

			//Convert AES Key Element to Bytes and store it in a file
			m := res.Bytes()
			functions.FileCrWr("../tmp/private/client/Result.txt", m)

			//Re-generate the 32bytes AES Key from the received key
			file, err := os.Open("../tmp/private/client/Result.txt") // For read access.
			functions.CheckError(err)
			data := bufio.NewReader(file)
    			test_key, err = data.Peek(32)
			functions.CheckError(err)
			}
			y()
			
			z := func() {
			defer timeTrack(time.Now(), "z")
			functions.AESdecrypt(test_key, nonce)
			}
			z()
			fmt.Println("Finished")
		} else {
			fmt.Println("Hash Verification Failed")
		}
	}
}
